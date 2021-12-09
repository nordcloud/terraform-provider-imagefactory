// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

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

func expandTemplateAwsConfig(in []interface{}) *graphql.NewTemplateAWSConfig {
	if len(in) == 0 {
		return nil
	}

	m := in[0].(map[string]interface{})
	return &graphql.NewTemplateAWSConfig{
		Region: graphql.String(m["region"].(string)),
	}
}

func expandTemplateAzureConfig(in []interface{}) *graphql.NewTemplateAZUREConfig {
	if len(in) == 0 {
		return nil
	}

	m := in[0].(map[string]interface{})

	e := graphql.Boolean(m["exclude_from_latest"].(bool))

	rr := []graphql.String{}
	for _, v := range m["replica_regions"].([]interface{}) {
		rr = append(rr, graphql.String(v.(string)))
	}

	out := &graphql.NewTemplateAZUREConfig{
		ExcludeFromLatest: &e,
		ReplicaRegions:    &rr,
	}

	return out
}

func expandTemplateConfig(in []interface{}) *graphql.NewTemplateConfig {
	if len(in) == 0 {
		return &graphql.NewTemplateConfig{}
	}

	m := in[0].(map[string]interface{})

	return &graphql.NewTemplateConfig{
		Aws:             expandTemplateAwsConfig(m["aws"].([]interface{})),
		Azure:           expandTemplateAzureConfig(m["azure"].([]interface{})),
		BuildComponents: expandTemplateComponents(m["build_components"].([]interface{})),
		TestComponents:  expandTemplateComponents(m["test_components"].([]interface{})),
		Notifications:   expandTemplateNotifications(m["notifications"].([]interface{})),
		Tags:            expandTemplateTags(m["tags"].([]interface{})),
	}
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
