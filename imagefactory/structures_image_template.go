package imagefactory

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
