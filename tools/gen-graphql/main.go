// gen-graphql generates pkg/graphql/graphql.go from GraphQL schema and operation files.
// Usage: go run ./tools/gen-graphql
//
// This tool replaces the abandoned graphql-codegen-golang npm package.
// It parses the GraphQL schema to extract types/inputs/enums and the operation
// files (*.graphql except schema.graphql) to extract query/mutation operations,
// then emits a single Go file with the same structure as the original generated file.

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

// ---------------------------------------------------------------------------
// Data model
// ---------------------------------------------------------------------------

type ScalarDef struct {
	Name   string
	GoType string
}

type EnumValue struct {
	ConstName string
	Value     string
}

type EnumDef struct {
	Name   string
	Values []EnumValue
}

type FieldDef struct {
	Name     string
	JSONName string
	GoType   string
	Optional bool   // pointer
	IsList   bool   // true if list type
	BaseType string // underlying GraphQL type name (e.g. "Account", "String")
}

type TypeDef struct {
	Name     string
	Fields   []FieldDef
	IsRootOp bool // true for Query/Mutation aggregate types
}

type OperationParam struct {
	Name    string
	GQLType string
	GoType  string
}

// SelectionNode represents a selection set (field or nested)
type SelectionNode struct {
	Name         string
	JSONName     string
	RawArgs      string // raw inline args text, e.g. "(input: {page: 0, limit: 1})"
	Children     []SelectionNode
	IsLeaf       bool
	IsList       bool // true if this field is a list type in the schema
	ListOptional bool // true if the list itself is nullable → *[]T, false → []T
	// For scalars: the Go type (string by default)
	GoType string
}

type OperationDef struct {
	Kind        string // "query" | "mutation"
	Name        string
	VarType     string // input variable Go type name (e.g. CustomerAccountIdInput)
	VarJSONName string
	VarGQLType  string
	HasVars     bool
	QueryStr    string
	// Root selection field and its tree
	RootField string
	RootJSON  string
	Selection []SelectionNode
}

// ---------------------------------------------------------------------------
// GraphQL schema parser
// ---------------------------------------------------------------------------

type Schema struct {
	Types    []TypeDef
	Inputs   []TypeDef
	Enums    []EnumDef
	Scalars  []ScalarDef
	Query    *TypeDef
	Mutation *TypeDef
}

// stripComments removes # line comments and """ block comments from graphql source
func stripComments(src string) string {
	// Remove """ ... """ block comments
	tripleRe := regexp.MustCompile(`"""[\s\S]*?"""`)
	src = tripleRe.ReplaceAllString(src, "")
	// Remove # line comments
	lines := strings.Split(src, "\n")
	for i, l := range lines {
		if idx := strings.Index(l, "#"); idx >= 0 {
			lines[i] = l[:idx]
		}
	}
	return strings.Join(lines, "\n")
}

// tokenize turns graphql text into tokens (words, braces, punctuation)
type tokenizer struct {
	tokens []string
	starts []int // byte offset of each token in src
	src    string
	pos    int
}

func tokenize(src string) *tokenizer {
	re := regexp.MustCompile(`[{}()\[\]!:,=@]|[^\s{}()\[\]!:,=@]+`)
	matches := re.FindAllStringIndex(src, -1)
	tokens := make([]string, len(matches))
	starts := make([]int, len(matches))
	for i, m := range matches {
		tokens[i] = src[m[0]:m[1]]
		starts[i] = m[0]
	}
	return &tokenizer{tokens: tokens, starts: starts, src: src}
}

func (t *tokenizer) peek() string {
	if t.pos >= len(t.tokens) {
		return ""
	}
	return t.tokens[t.pos]
}

func (t *tokenizer) next() string {
	v := t.peek()
	t.pos++
	return v
}

func (t *tokenizer) expect(v string) {
	got := t.next()
	if got != v {
		panic(fmt.Sprintf("expected %q got %q", v, got))
	}
}

func (t *tokenizer) skipBlock() {
	depth := 1
	for depth > 0 && t.pos < len(t.tokens) {
		tok := t.next()
		switch tok {
		case "{":
			depth++
		case "}":
			depth--
		}
	}
}

// parseGQLType reads a GraphQL type like String!, [String!]!, Input etc.
// Returns the base name (without ! and []) and whether it's optional.
func parseGQLType(t *tokenizer) (base string, list bool, optional bool) {
	optional = true
	if t.peek() == "[" {
		t.next() // consume [
		base, _, _ = parseGQLType(t)
		t.expect("]")
		list = true
	} else {
		base = t.next()
	}
	if t.peek() == "!" {
		t.next()
		optional = false
	}
	return
}

// gqlTypeToGo converts a GraphQL type to a Go type string.
// optional=true means the field can be absent → pointer.
// Lists are always pointers (consistent with original generator behavior).
func gqlTypeToGo(base string, list bool, optional bool) string {
	goBase := scalarGoType(base)
	if list {
		// Lists are always *[]T regardless of nullability
		return "*[]" + goBase
	}
	if optional {
		return "*" + goBase
	}
	return goBase
}

func scalarGoType(name string) string {
	switch name {
	case "String", "ID":
		return "String"
	case "Int":
		return "Int"
	case "Float":
		return "Float"
	case "Boolean":
		return "Boolean"
	default:
		return name
	}
}

func parseSchema(src string) Schema {
	src = stripComments(src)
	t := tokenize(src)

	var schema Schema
	// Collect Query and Mutation fields from base type + all extend type blocks
	queryFields := []FieldDef{}
	mutationFields := []FieldDef{}

	for t.pos < len(t.tokens) {
		tok := t.peek()
		switch tok {
		case "type", "input":
			kind := t.next()
			name := t.next()
			// skip implements / directives before {
			for t.peek() != "{" && t.peek() != "" {
				t.next()
			}
			if t.peek() != "{" {
				continue
			}
			t.expect("{")
			fields := parseFieldBlock(t)
			td := TypeDef{Name: name, Fields: fields}
			if kind == "input" {
				schema.Inputs = append(schema.Inputs, td)
			} else if name == "Query" {
				queryFields = append(queryFields, fields...)
			} else if name == "Mutation" {
				mutationFields = append(mutationFields, fields...)
			} else {
				schema.Types = append(schema.Types, td)
			}
		case "extend":
			t.next()
			inner := t.peek()
			switch inner {
			case "type":
				t.next()
				extName := t.next() // e.g. Query, Mutation
				if t.peek() == "{" {
					t.expect("{")
					fields := parseFieldBlock(t)
					switch extName {
					case "Query":
						queryFields = append(queryFields, fields...)
					case "Mutation":
						mutationFields = append(mutationFields, fields...)
					}
					// for other extend type (unlikely), skip
				}
			case "input":
				t.next()
				t.next() // name
				if t.peek() == "{" {
					t.expect("{")
					t.skipBlock()
				}
			default:
				t.next()
			}
		case "enum":
			t.next()
			name := t.next()
			t.expect("{")
			var values []EnumValue
			for t.peek() != "}" && t.peek() != "" {
				val := t.next()
				if val == "}" {
					break
				}
				constName := enumConstName(name, val)
				values = append(values, EnumValue{ConstName: constName, Value: val})
			}
			t.expect("}")
			schema.Enums = append(schema.Enums, EnumDef{Name: name, Values: values})
		case "scalar":
			t.next()
			name := t.next()
			// custom scalars not used here — we map them to String
			schema.Scalars = append(schema.Scalars, ScalarDef{Name: name, GoType: "String"})
		case "schema":
			t.next()
			if t.peek() == "{" {
				t.expect("{")
				t.skipBlock()
			}
		default:
			t.next()
		}
	}

	if len(queryFields) > 0 {
		td := TypeDef{Name: "Query", Fields: queryFields, IsRootOp: true}
		schema.Query = &td
	}
	if len(mutationFields) > 0 {
		td := TypeDef{Name: "Mutation", Fields: mutationFields, IsRootOp: true}
		schema.Mutation = &td
	}

	return schema
}

func parseFieldBlock(t *tokenizer) []FieldDef {
	var fields []FieldDef
	for t.peek() != "}" && t.peek() != "" {
		fname := t.next()
		if fname == "}" {
			break
		}
		// skip field arguments (e.g. accounts(input: CustomerAccountsResolverInput!): AccountResults)
		if t.peek() == "(" {
			t.next()
			depth := 1
			for depth > 0 && t.peek() != "" {
				tok := t.next()
				switch tok {
				case "(":
					depth++
				case ")":
					depth--
				}
			}
		}
		if t.peek() != ":" {
			continue
		}
		t.expect(":")
		base, list, optional := parseGQLType(t)
		// skip default value
		if t.peek() == "=" {
			t.next()
			t.next()
		}
		goType := gqlTypeToGo(base, list, optional)
		jsonTag := fname
		goName := goFieldName(fname)
		fields = append(fields, FieldDef{
			Name:     goName,
			JSONName: jsonTag,
			GoType:   goType,
			Optional: optional,
			IsList:   list,
			BaseType: base,
		})
	}
	if t.peek() == "}" {
		t.next()
	}
	return fields
}

// ---------------------------------------------------------------------------
// Operation file parser
// ---------------------------------------------------------------------------

// typeFieldInfo holds per-field metadata from the schema.
type typeFieldInfo struct {
	IsList   bool
	Optional bool // for lists: whether the list itself is nullable
	BaseType string
}

// typeMap maps typeName → (jsonFieldName → info)
type typeMap map[string]map[string]typeFieldInfo

// buildTypeMap builds a lookup map from the parsed schema so we can resolve
// whether a selection-set field is a list when generating Go structs.
func buildTypeMap(schema Schema) typeMap {
	tm := make(typeMap)
	allTypes := append(schema.Types, schema.Inputs...)
	if schema.Query != nil {
		allTypes = append(allTypes, *schema.Query)
	}
	if schema.Mutation != nil {
		allTypes = append(allTypes, *schema.Mutation)
	}
	for _, td := range allTypes {
		m := make(map[string]typeFieldInfo)
		for _, f := range td.Fields {
			m[f.JSONName] = typeFieldInfo{IsList: f.IsList, Optional: f.Optional, BaseType: f.BaseType}
		}
		tm[td.Name] = m
	}
	return tm
}

// annotateSelectionList walks the selection tree and marks IsList based on
// the schema type map. parentType is the GraphQL type name of the parent object.
func annotateSelectionList(nodes []SelectionNode, parentType string, tm typeMap) []SelectionNode {
	fm := tm[parentType]
	result := make([]SelectionNode, len(nodes))
	for i, n := range nodes {
		info, ok := fm[n.JSONName]
		if ok {
			n.IsList = info.IsList
			n.ListOptional = info.Optional
		}
		// Recurse into children using the child's base type
		if len(n.Children) > 0 && ok && info.BaseType != "" {
			n.Children = annotateSelectionList(n.Children, info.BaseType, tm)
		}
		result[i] = n
	}
	return result
}

func parseOperations(src string, tm typeMap, queryRootType, mutationRootType map[string]string) []OperationDef {
	src = stripComments(src)
	t := tokenize(src)

	var ops []OperationDef

	for t.pos < len(t.tokens) {
		tok := t.peek()
		if tok != "query" && tok != "mutation" {
			t.next()
			continue
		}
		kind := t.next()
		name := t.next()

		var varName, varGQLType string
		hasVars := false
		if t.peek() == "(" {
			t.next() // (
			t.next() // $varName
			t.expect(":")
			base, list, optional := parseGQLType(t)
			varGQLType = base
			if list {
				varGQLType = "[" + base + "]"
			}
			if !optional {
				varGQLType += "!"
			}
			// consume )
			for t.peek() != ")" && t.peek() != "" {
				t.next()
			}
			t.expect(")")
			varName = base // the GraphQL input type name
			_ = varName
			hasVars = true
		}

		if t.peek() != "{" {
			continue
		}
		t.expect("{")

		// Parse root field
		rootField := t.next()
		rootJSON := rootField

		// Parse root args if any (inline literal args like images(input: { page: 0, limit: 1 }))
		if t.peek() == "(" {
			t.next()
			depth := 1
			for depth > 0 && t.peek() != "" {
				tok := t.next()
				switch tok {
				case "(":
					depth++
				case ")":
					depth--
				}
			}
		}

		var selection []SelectionNode
		if t.peek() == "{" {
			t.expect("{")
			selection = parseSelectionSet(t)
		}

		// consume outer }
		if t.peek() == "}" {
			t.next()
		}

		// Determine the return type of the root field from Query/Mutation type map
		var rootReturnType string
		if kind == "query" {
			rootReturnType = queryRootType[rootJSON]
		} else {
			rootReturnType = mutationRootType[rootJSON]
		}
		// Annotate selection nodes with IsList info from schema
		if rootReturnType != "" && len(selection) > 0 {
			selection = annotateSelectionList(selection, rootReturnType, tm)
		}

		// Build the query string (after annotation, using original selection tree)
		queryStr := buildQueryString(kind, name, hasVars, varGQLType, rootField, selection)

		// Determine Go variable type from the gql type
		goVarType := ""
		varJSONName := "input"
		if hasVars {
			// Extract base type from varGQLType (strip !, [])
			base := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSuffix(varGQLType, "!"), "["), "]")
			base = strings.TrimSuffix(base, "!")
			goVarType = base
		}

		ops = append(ops, OperationDef{
			Kind:        kind,
			Name:        name,
			VarType:     goVarType,
			VarJSONName: varJSONName,
			VarGQLType:  varGQLType,
			HasVars:     hasVars,
			QueryStr:    queryStr,
			RootField:   goFieldName(rootField),
			RootJSON:    rootJSON,
			Selection:   selection,
		})
	}
	return ops
}

func parseSelectionSet(t *tokenizer) []SelectionNode {
	var nodes []SelectionNode
	for t.peek() != "}" && t.peek() != "" {
		fname := t.next()
		if fname == "}" {
			break
		}
		// inline args — capture raw text from source to preserve formatting
		rawArgs := ""
		if t.peek() == "(" {
			startIdx := t.starts[t.pos] // byte offset of "("
			t.next()                    // consume (
			depth := 1
			for depth > 0 && t.peek() != "" {
				tok := t.next()
				switch tok {
				case "(":
					depth++
				case ")":
					depth--
				}
			}
			// end is the position just after the closing ")"
			endIdx := t.starts[t.pos-1] + 1 // +1 to include ")"
			rawArgs = t.src[startIdx:endIdx]
		}

		node := SelectionNode{
			Name:     fname,
			JSONName: fname,
			RawArgs:  rawArgs,
		}
		if t.peek() == "{" {
			t.expect("{")
			node.Children = parseSelectionSet(t)
			node.IsLeaf = false
		} else {
			node.IsLeaf = true
			node.GoType = "string"
		}
		nodes = append(nodes, node)
	}
	if t.peek() == "}" {
		t.next()
	}
	return nodes
}

// ---------------------------------------------------------------------------
// Query string builder (reproduces the .graphql operation format)
// ---------------------------------------------------------------------------

func buildQueryString(kind, name string, hasVars bool, varGQLType, rootField string, sel []SelectionNode) string {
	var b strings.Builder
	b.WriteString(kind)
	b.WriteString(" ")
	b.WriteString(name)
	if hasVars {
		b.WriteString("($input: ")
		b.WriteString(varGQLType)
		b.WriteString(")")
	}
	b.WriteString(" {\n")
	b.WriteString("  ")
	b.WriteString(rootField)
	if hasVars {
		b.WriteString("(input: $input)")
	}
	if len(sel) > 0 {
		b.WriteString(" {\n")
		writeSelectionSet(&b, sel, 4)
		b.WriteString("  }")
	}
	b.WriteString("\n}")
	return b.String()
}

func writeSelectionSet(b *strings.Builder, nodes []SelectionNode, indent int) {
	pad := strings.Repeat(" ", indent)
	for _, n := range nodes {
		b.WriteString(pad)
		b.WriteString(n.Name)
		if n.RawArgs != "" {
			b.WriteString(n.RawArgs)
		}
		if len(n.Children) > 0 {
			b.WriteString(" {\n")
			writeSelectionSet(b, n.Children, indent+2)
			b.WriteString(pad)
			b.WriteString("}")
		}
		b.WriteString("\n")
	}
}

// ---------------------------------------------------------------------------
// Response struct builder from selection set
// ---------------------------------------------------------------------------

// selectionToGoStructFull builds an inline Go struct from a selection set.
// Fields named "results", or any field with IsList=true, are wrapped as *[]struct{...}.
func selectionToGoStructFull(nodes []SelectionNode, indent int) string {
	var b strings.Builder
	pad := strings.Repeat("\t", indent)
	for _, n := range nodes {
		fname := goFieldName(n.Name)
		jsonTag := n.JSONName
		if len(n.Children) == 0 {
			b.WriteString(fmt.Sprintf("%s%s string `json:\"%s\"`\n", pad, fname, jsonTag))
		} else {
			isList := n.IsList || n.Name == "results"
			if isList {
				listType := "*[]"
				if !n.ListOptional {
					listType = "[]"
				}
				b.WriteString(fmt.Sprintf("%s%s %sstruct {\n", pad, fname, listType))
				b.WriteString(selectionToGoStructFull(n.Children, indent+1))
				b.WriteString(fmt.Sprintf("%s} `json:\"%s\"`\n", pad, jsonTag))
			} else {
				b.WriteString(fmt.Sprintf("%s%s struct {\n", pad, fname))
				b.WriteString(selectionToGoStructFull(n.Children, indent+1))
				b.WriteString(fmt.Sprintf("%s} `json:\"%s\"`\n", pad, jsonTag))
			}
		}
	}
	return b.String()
}

// ---------------------------------------------------------------------------
// Name helpers
// ---------------------------------------------------------------------------

// enumConstName produces Go const name like AccountStatusACCESSSUCCESS or EBSVolumeTypeGp2
func enumConstName(typeName, value string) string {
	// Check if value is lowercase (e.g. "gp2", "gp3") - preserve original casing with uppercase first letter
	hasLower := false
	for _, r := range value {
		if unicode.IsLower(r) {
			hasLower = true
			break
		}
	}

	var sb strings.Builder
	sb.WriteString(typeName)

	if hasLower {
		// Mixed or lowercase value: title-case first letter, rest as-is, strip underscores
		parts := strings.Split(value, "_")
		for _, p := range parts {
			if len(p) == 0 {
				continue
			}
			r := []rune(p)
			r[0] = unicode.ToUpper(r[0])
			sb.WriteString(string(r))
		}
	} else {
		// All-uppercase value: strip underscores, keep all uppercase
		parts := strings.Split(value, "_")
		for _, p := range parts {
			sb.WriteString(p)
		}
	}
	return sb.String()
}

// goFieldName converts camelCase/lower field to Go exported field name
func goFieldName(name string) string {
	if name == "" {
		return ""
	}
	// Handle fields with underscores (e.g. AWS_ACCESS_KEY_ID → AWSACCESSKEYID)
	if strings.Contains(name, "_") {
		parts := strings.Split(name, "_")
		var sb strings.Builder
		for _, p := range parts {
			sb.WriteString(p)
		}
		return sb.String()
	}
	// Handle all-caps abbreviations at start: id → ID, etc.
	specialMap := map[string]string{
		"id": "ID",
	}
	if v, ok := specialMap[strings.ToLower(name)]; ok && strings.ToLower(name) == name {
		return v
	}
	r := []rune(name)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// ---------------------------------------------------------------------------
// Code generator
// ---------------------------------------------------------------------------

const fileHeader = `package graphql

// Code generated by tools/gen-graphql ; DO NOT EDIT.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	*http.Client
	Url string
}

// NewClient creates a GraphQL client ready to use.
func NewClient(url string) *Client {
	return &Client{
		Client: &http.Client{},
		Url:    url,
	}
}

type GraphQLOperation struct {
	Query         string          ` + "`" + `json:"query"` + "`" + `
	OperationName string          ` + "`" + `json:"operationName,omitempty"` + "`" + `
	Variables     json.RawMessage ` + "`" + `json:"variables,omitempty"` + "`" + `
}

type GraphQLResponse struct {
	Data   json.RawMessage ` + "`" + `json:"data,omitempty"` + "`" + `
	Errors []GraphQLError  ` + "`" + `json:"errors,omitempty"` + "`" + `
}

type GraphQLError map[string]interface{}

func (err GraphQLError) Error() string {
	return fmt.Sprintf("graphql: %v", map[string]interface{}(err))
}

func (resp *GraphQLResponse) Error() string {
	if len(resp.Errors) == 0 {
		return ""
	}
	errs := strings.Builder{}
	for _, err := range resp.Errors {
		errs.WriteString(err.Error())
		errs.WriteString("\n")
	}
	return errs.String()
}

func execute(client *http.Client, req *http.Request) (*GraphQLResponse, error) {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return unmarshalGraphQLReponse(body)
}

func unmarshalGraphQLReponse(b []byte) (*GraphQLResponse, error) {
	resp := GraphQLResponse{}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	if len(resp.Errors) > 0 {
		return &resp, &resp
	}
	return &resp, nil
}

`

const scalarsSection = `//
// Scalars
//

type Int int32
type Float float64
type Boolean bool
type String string
type ID string

`

// operationTmpl generates code for a single operation
const operationTmpl = `//
// {{.KindComment}} {{.Name}}{{if .HasVars}}({{.VarGQLName}}: {{.VarGQLType}}){{end}}
//

{{- if .HasVars}}
type {{.Name}}Variables struct {
	Input {{.VarType}} ` + "`" + `json:"input"` + "`" + `
}
{{- else}}

{{- end}}

type {{.Name}}Response struct {
	{{.RootFieldName}} {{.RootGoType}}
}

type {{.Name}}Request struct {
	*http.Request
}

func New{{.Name}}Request(url string{{if .HasVars}}, vars *{{.Name}}Variables{{end}}) (*{{.Name}}Request, error) {
{{- if .HasVars}}
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: ` + "`" + `{{.QueryStr}}` + "`" + `,
	})
{{- else}}
	b, err := json.Marshal(&GraphQLOperation{
		Query: ` + "`" + `{{.QueryStr}}` + "`" + `,
	})
{{- end}}
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &{{.Name}}Request{req}, nil
}

func (req *{{.Name}}Request) Execute(client *http.Client) (*{{.Name}}Response, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result {{.Name}}Response
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func {{.Name}}(url string, client *http.Client{{if .HasVars}}, vars *{{.Name}}Variables{{end}}) (*{{.Name}}Response, error) {
	req, err := New{{.Name}}Request(url{{if .HasVars}}, vars{{end}})
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) {{.Name}}({{if .HasVars}}vars *{{.Name}}Variables{{end}}) (*{{.Name}}Response, error) {
	return {{.Name}}(client.Url, client.Client{{if .HasVars}}, vars{{end}})
}
`

type opTmplData struct {
	KindComment   string
	Name          string
	HasVars       bool
	VarGQLName    string
	VarGQLType    string
	VarType       string
	QueryStr      string
	RootFieldName string
	RootGoType    string
}

// buildRootGoType generates the inline struct type for the response root field
func buildRootGoType(rootJSON string, sel []SelectionNode, isDelete bool) string {
	if isDelete {
		return "string `json:\"" + rootJSON + "\"`"
	}
	if len(sel) == 0 {
		// scalar return
		return "struct{} `json:\"" + rootJSON + "\"`"
	}
	var b strings.Builder
	b.WriteString("struct {\n")
	b.WriteString(selectionToGoStructFull(sel, 1))
	b.WriteString("} `json:\"")
	b.WriteString(rootJSON)
	b.WriteString("\"`")
	return b.String()
}

func generateOperations(ops []OperationDef) string {
	tmpl := template.Must(template.New("op").Parse(operationTmpl))
	var buf bytes.Buffer
	for _, op := range ops {
		isDelete := strings.HasPrefix(op.Name, "Delete") || strings.HasPrefix(op.Name, "Recheck")
		// "Recheck" returns Account not bool, so only Delete* are bool
		isDelete = strings.HasPrefix(op.Name, "Delete")

		rootGoType := buildRootGoType(op.RootJSON, op.Selection, isDelete)

		kindComment := "query"
		if op.Kind == "mutation" {
			kindComment = "mutation"
		}

		data := opTmplData{
			KindComment:   kindComment,
			Name:          op.Name,
			HasVars:       op.HasVars,
			VarGQLName:    "$input",
			VarGQLType:    op.VarGQLType,
			VarType:       op.VarType,
			QueryStr:      op.QueryStr,
			RootFieldName: goFieldName(op.RootJSON),
			RootGoType:    rootGoType,
		}
		if err := tmpl.Execute(&buf, data); err != nil {
			panic(err)
		}
	}
	return buf.String()
}

func generateEnums(enums []EnumDef) string {
	var b strings.Builder
	b.WriteString("//\n// Enums\n//\n\n")
	// Sort by name
	sort.Slice(enums, func(i, j int) bool { return enums[i].Name < enums[j].Name })
	for _, e := range enums {
		b.WriteString(fmt.Sprintf("type %s string\n\n", e.Name))
		b.WriteString("const (\n")
		for _, v := range e.Values {
			b.WriteString(fmt.Sprintf("\t%s %s = %q\n", v.ConstName, e.Name, v.Value))
		}
		b.WriteString(")\n\n")
	}
	return b.String()
}

func generateTypes(name string, types []TypeDef) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("//\n// %s\n//\n\n", name))
	sort.Slice(types, func(i, j int) bool { return types[i].Name < types[j].Name })
	for _, t := range types {
		b.WriteString(generateSingleType(t))
	}
	return b.String()
}

func generateSingleType(t TypeDef) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("type %s struct {\n", t.Name))
	// Sort fields by name for deterministic output
	fields := make([]FieldDef, len(t.Fields))
	copy(fields, t.Fields)
	sort.Slice(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
	for _, f := range fields {
		b.WriteString(fmt.Sprintf("\t%s %s `json:\"%s", f.Name, f.GoType, f.JSONName))
		if strings.HasPrefix(f.GoType, "*") {
			b.WriteString(",omitempty")
		}
		b.WriteString("\"`\n")
	}
	b.WriteString("}\n\n")
	return b.String()
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: gen-graphql <schema.graphql> <output.go>\n")
		os.Exit(1)
	}
	schemaPath := os.Args[1]
	outPath := os.Args[2]

	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read schema: %v\n", err)
		os.Exit(1)
	}

	schema := parseSchema(string(schemaBytes))

	// Build type map for selection annotation
	tm := buildTypeMap(schema)

	// Build root field → return type maps for Query and Mutation
	queryRootType := make(map[string]string)
	mutationRootType := make(map[string]string)
	if schema.Query != nil {
		for _, f := range schema.Query.Fields {
			queryRootType[f.JSONName] = f.BaseType
		}
	}
	if schema.Mutation != nil {
		for _, f := range schema.Mutation.Fields {
			mutationRootType[f.JSONName] = f.BaseType
		}
	}

	// Operation files: all *.graphql in the same directory as schema, except schema itself
	schemaDir := filepath.Dir(schemaPath)
	entries, err := filepath.Glob(filepath.Join(schemaDir, "*.graphql"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "glob error: %v\n", err)
		os.Exit(1)
	}

	var allOps []OperationDef
	sort.Strings(entries)
	for _, path := range entries {
		if filepath.Base(path) == filepath.Base(schemaPath) {
			continue
		}
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot read %s: %v\n", path, err)
			os.Exit(1)
		}
		ops := parseOperations(string(data), tm, queryRootType, mutationRootType)
		allOps = append(allOps, ops...)
	}

	var buf bytes.Buffer
	buf.WriteString(fileHeader)

	// Operations
	buf.WriteString(generateOperations(allOps))

	// Scalars
	buf.WriteString(scalarsSection)

	// Enums
	buf.WriteString(generateEnums(schema.Enums))

	// Inputs
	buf.WriteString(generateTypes("Inputs", schema.Inputs))

	// Objects
	buf.WriteString(generateTypes("Objects", schema.Types))

	// Root operation types (Query, Mutation) - needed by sdk/client.go
	if schema.Mutation != nil {
		buf.WriteString(generateSingleType(*schema.Mutation))
	}
	if schema.Query != nil {
		buf.WriteString(generateSingleType(*schema.Query))
	}

	if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "cannot write output: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated %s\n", outPath)
	fmt.Printf("  %d operations, %d types, %d inputs, %d enums\n",
		len(allOps), len(schema.Types), len(schema.Inputs), len(schema.Enums))
}
