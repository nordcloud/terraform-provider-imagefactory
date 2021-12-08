// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func expandTemplateComponents(in []interface{}) *[]graphql.NewTemplateComponent {
	out := []graphql.NewTemplateComponent{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewTemplateComponent{
			ID: graphql.String(m["id"].(string)),
		})
	}

	return &out
}

func expandTemplateNotifications(in []interface{}) *[]graphql.NewNotification {
	out := []graphql.NewNotification{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewNotification{
			Type: graphql.NotificationType(m["type"].(string)),
			Uri:  graphql.String(m["uri"].(string)),
		})
	}

	return &out
}

func expandTemplateTags(in []interface{}) *[]graphql.NewTag {
	out := []graphql.NewTag{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewTag{
			Key:   graphql.String(m["key"].(string)),
			Value: graphql.String(m["value"].(string)),
		})
	}

	return &out
}

func expandTemplateConfig(in []interface{}) *graphql.NewTemplateConfig {
	if len(in) == 0 {
		return &graphql.NewTemplateConfig{}
	}

	m := in[0].(map[string]interface{})

	out := &graphql.NewTemplateConfig{
		BuildComponents: expandTemplateComponents(m["build_components"].([]interface{})),
		TestComponents:  expandTemplateComponents(m["test_components"].([]interface{})),
		Notifications:   expandTemplateNotifications(m["notifications"].([]interface{})),
		Tags:            expandTemplateTags(m["tags"].([]interface{})),
	}

	awsConfig := m["aws"].([]interface{})
	if len(awsConfig) > 0 {
		t := awsConfig[0].(map[string]interface{})
		out.Aws = &graphql.NewTemplateAWSConfig{
			Region: graphql.String(t["region"].(string)),
		}
	}

	return out
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
