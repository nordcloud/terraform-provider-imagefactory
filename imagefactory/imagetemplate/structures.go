// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import (
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func expandTemplateComponents(in []interface{}) *[]graphql.NewTemplateComponent {
	components := []graphql.NewTemplateComponent{}
	for i := range in {
		bc := in[i].(map[string]interface{})

		components = append(components, graphql.NewTemplateComponent{
			ID: graphql.String(bc["id"].(string)),
		})
	}

	return &components
}

func expandTemplateConfig(in []interface{}) *graphql.NewTemplateConfig {
	templateConfig := &graphql.NewTemplateConfig{}

	if len(in) == 0 {
		return templateConfig
	}

	m := in[0].(map[string]interface{})

	awsConfig := m["aws"].([]interface{})
	if len(awsConfig) > 0 {
		t := awsConfig[0].(map[string]interface{})
		templateConfig.Aws = &graphql.NewTemplateAWSConfig{
			Region: graphql.String(t["region"].(string)),
		}
	}

	templateConfig.BuildComponents = expandTemplateComponents(m["build_components"].([]interface{}))
	templateConfig.TestComponents = expandTemplateComponents(m["test_components"].([]interface{}))

	return templateConfig
}

func flattenTemplateState(in graphql.TemplateState) map[string]string {
	out := map[string]string{
		"status": string(in.Status),
	}
	if in.Error != nil {
		out["error"] = string(*in.Error)
	}

	return out
}
