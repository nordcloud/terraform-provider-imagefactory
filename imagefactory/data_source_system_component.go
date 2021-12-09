// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	},
	"cloud_provider": {
		Type:     schema.TypeString,
		Required: true,
	},
}

func dataSourceSystemComponent() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSystemComponentRead,
		Schema:      componentSchema,
	}
}

func dataSourceSystemComponentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	component, err := config.client.GetSystemComponent(d.Get("name").(string), d.Get("cloud_provider").(string), d.Get("stage").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(component.ID))

	return diags
}
