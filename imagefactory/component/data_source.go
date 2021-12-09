// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: systemComponentRead,
		Schema:      componentSchema,
	}
}

func systemComponentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*config.Config)

	component, err := config.APIClient.GetSystemComponent(d.Get("name").(string), d.Get("cloud_provider").(string), d.Get("stage").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(component.ID))
	if err := d.Set("name", component.Name); err != nil {
		return diag.FromErr(err)
	}
	// if err := d.Set("cloud_provider", component.Providers); err != nil {
	// 	return diag.FromErr(err)
	// }
	if err := d.Set("stage", component.Stage); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
