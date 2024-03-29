// Copyright 2021-2024 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var systemComponentSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var customComponentSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"stage": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validStages, false),
	},
	"cloud_provider": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validCloudProviders, false),
	},
}

var contentComponentResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"script": {
			Type:     schema.TypeString,
			Required: true,
		},
		"provisioner": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(validProvisioners, false),
		},
	},
}

var componentSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"stage": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validStages, false),
	},
	"cloud_providers": {
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validCloudProviders, false),
		},
	},
	"os_types": {
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validOSTypes, false),
		},
	},
	"content": {
		Type:     schema.TypeList,
		Required: true,
		Elem:     contentComponentResource,
	},
	"rebuild_templates": {
		Type:     schema.TypeBool,
		Optional: true,
		Description: "Trigger rebuild of templates using this component. " +
			"Only applicable when content is updated. Default is false.",
	},
}
