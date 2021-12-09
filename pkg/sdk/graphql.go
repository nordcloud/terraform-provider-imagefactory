package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GraphQLAPI interface {
	Execute(req *http.Request, result interface{}) error
}

type GraphQLClient struct {
	client  *http.Client
	headers map[string]string
}

func NewGraphQLClient(httpClient *http.Client, headers map[string]string) GraphQLAPI {
	return &GraphQLClient{
		&http.Client{},
		headers,
	}
}

func (g *GraphQLClient) Execute(req *http.Request, result interface{}) error {
	for k, v := range g.headers {
		req.Header.Add(k, v)
	}

	resp, err := g.execute(req)
	if err != nil {
		return err
	}

	if len(resp.Errors) != 0 {
		return fmt.Errorf("query to %s failed: %s", req.URL.Path, resp.Error())
	}

	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return fmt.Errorf("failed to unmarshal response JSON '%s': %s", resp.Data, err)
	}

	return nil
}

func (g *GraphQLClient) execute(req *http.Request) (*GraphQLResponse, error) {
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
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

type GraphQLResponse struct {
	Data   json.RawMessage `json:"data,omitempty"`
	Errors []GraphQLError  `json:"errors,omitempty"`
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
