// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var componentSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"stage": {
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringInSlice(validStages, false),
	},
	"cloud_provider": {
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringInSlice(validCloudProviders, false),
	},
}
