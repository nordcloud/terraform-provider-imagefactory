// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package role

import (
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func expandRoleRules(in []interface{}) *[]graphql.NewRule {
	out := []graphql.NewRule{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewRule{
			Actions:   expandActions(m["actions"].([]interface{})),
			Resources: expandResources(m["resources"].([]interface{})),
		})
	}

	return &out
}

func expandActions(in []interface{}) *[]graphql.Action {
	out := []graphql.Action{}
	for _, v := range in {
		out = append(out, graphql.Action(v.(string)))
	}

	return &out
}

func expandResources(in []interface{}) *[]graphql.Resource {
	out := []graphql.Resource{}
	for _, v := range in {
		out = append(out, graphql.Resource(v.(string)))
	}

	return &out
}

func flattenRoleRules(in *[]graphql.Rule) []interface{} {
	if in == nil {
		return make([]interface{}, 0)
	}

	out := make([]interface{}, len(*in))
	for i, rule := range *in {
		oi := make(map[string]interface{})

		oi["actions"] = flattenActions(rule.Actions)
		oi["resources"] = flattenResources(rule.Resources)

		out[i] = oi
	}

	return out
}

func flattenActions(in *[]graphql.Action) []string {
	out := []string{}

	if in == nil {
		return out
	}

	for _, v := range *in {
		out = append(out, string(v))
	}

	return out
}

func flattenResources(in *[]graphql.Resource) []string {
	out := []string{}

	if in == nil {
		return out
	}

	for _, v := range *in {
		out = append(out, string(v))
	}

	return out
}
