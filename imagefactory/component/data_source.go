// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
)

func DataSourceSystem() *schema.Resource {
	return &schema.Resource{
		ReadContext: systemComponentRead,
		Schema:      systemComponentSchema,
	}
}

func DataSourceCustom() *schema.Resource {
	return &schema.Resource{
		ReadContext: customComponentRead,
		Schema:      customComponentSchema,
	}
}

func systemComponentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	component, err := c.APIClient.GetSystemComponent(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(component.ID))
	if err := d.Set("name", component.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func customComponentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	component, err := c.APIClient.GetCustomComponent(d.Get("name").(string), d.Get("cloud_provider").(string), d.Get("stage").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(component.ID))
	if err := d.Set("name", component.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
