package imagetemplate

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var templateComponentResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var awsTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"region": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var templateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"test_components": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateComponentResource,
		},
		"build_components": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateComponentResource,
		},
		"aws": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     awsTemplateConfigResource,
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
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringInSlice([]string{
			"AWS",
			"AZURE",
			"GCP",
			"IBMCLOUD",
			"VMWARE",
		}, false),
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
