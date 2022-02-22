// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var awsTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"region": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

var azureTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"exclude_from_latest": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"replica_regions": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(validAzureRegions, false),
			},
		},
	},
}

var templateComponentResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var templateNotificationsResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(validNotificationTypes, false),
		},
		"uri": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var templateTagsResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var templateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"aws": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     awsTemplateConfigResource,
		},
		"azure": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     azureTemplateConfigResource,
		},
		"build_components": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateComponentResource,
		},
		"test_components": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateComponentResource,
		},
		"notifications": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateNotificationsResource,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateTagsResource,
		},
		"scope": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(validScopes, false),
			Default:      graphql.ScopePUBLIC,
		},
	},
}

var templateStateSchema = &schema.Schema{
	Type: schema.TypeString,
	Elem: map[string]*schema.Schema{
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"error": {
			Type:     schema.TypeString,
			Computed: true,
		},
	},
}

var templateSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"distribution_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"cloud_provider": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validCloudProviders, false),
	},
	"config": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     templateConfigResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     templateStateSchema,
	},
}
