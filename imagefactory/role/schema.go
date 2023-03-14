// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package role

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var roleRulesResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"actions": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(validRuleActions, false),
			},
		},
		"resources": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(validResourceActions, false),
			},
		},
	},
}

var roleSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"rules": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     roleRulesResource,
	},
}
