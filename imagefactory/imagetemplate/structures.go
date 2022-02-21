// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import (
	"errors"

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

func expandTemplateAwsConfig(in []interface{}, scope graphql.Scope) (*graphql.NewTemplateAWSConfig, error) {
	if len(in) == 0 || scope == graphql.ScopeCHINA {
		return nil, nil
	}

	m := in[0].(map[string]interface{})
	if m["region"] == nil || m["region"].(string) == "" {
		return nil, errors.New("AWS regions is required for the AWS template with PUBLIC scope.")
	}

	region := graphql.String(m["region"].(string))
	return &graphql.NewTemplateAWSConfig{
		Region: &region,
	}, nil
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

func expandTemplateConfig(in []interface{}) (*graphql.NewTemplateConfig, error) {
	if len(in) == 0 {
		return &graphql.NewTemplateConfig{}, nil
	}

	m := in[0].(map[string]interface{})
	scope := m["scope"].(graphql.Scope)

	awsCfg, err := expandTemplateAwsConfig(m["aws"].([]interface{}), scope)
	if err != nil {
		return nil, err
	}

	return &graphql.NewTemplateConfig{
		Aws:             awsCfg,
		Azure:           expandTemplateAzureConfig(m["azure"].([]interface{})),
		BuildComponents: expandTemplateComponents(m["build_components"].([]interface{})),
		TestComponents:  expandTemplateComponents(m["test_components"].([]interface{})),
		Notifications:   expandTemplateNotifications(m["notifications"].([]interface{})),
		Tags:            expandTemplateTags(m["tags"].([]interface{})),
		Scope:           &scope,
	}, nil
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
