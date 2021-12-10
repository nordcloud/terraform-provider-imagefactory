// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

import (
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func expandOSTypes(in []interface{}) *[]graphql.OSType {
	out := []graphql.OSType{}
	for _, v := range in {
		out = append(out, graphql.OSType(v.(string)))
	}

	return &out
}

func expandProviders(in []interface{}) *[]graphql.Provider {
	out := []graphql.Provider{}
	for _, v := range in {
		out = append(out, graphql.Provider(v.(string)))
	}

	return &out
}

func expandContent(in []interface{}) graphql.NewVersionedContent {
	if len(in) == 0 {
		return graphql.NewVersionedContent{}
	}

	m := in[0].(map[string]interface{})
	return graphql.NewVersionedContent{
		Script:            graphql.String(m["script"].(string)),
		ScriptProvisioner: graphql.ShellScriptProvisioner(m["provisioner"].(string)),
	}
}

func flattenOSTypes(in *[]graphql.OSType) []string {
	out := []string{}

	if in == nil {
		return out
	}

	for _, v := range *in {
		out = append(out, string(v))
	}

	return out
}

func flattenProviders(in *[]graphql.Provider) []string {
	out := []string{}

	if in == nil {
		return out
	}

	for _, v := range *in {
		out = append(out, string(v))
	}

	return out
}
