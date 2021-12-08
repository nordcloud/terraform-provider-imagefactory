// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"context"
	"strconv"
	"time"

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
	"type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"stage": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cloud_providers": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"os_types": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"content": {
		Type:     schema.TypeList,
		Computed: true,
		Elem:     contentSchema,
	},
}

var contentSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"version": {
			Type:     schema.TypeString,
			Computed: true,
		},
	},
}

func dataSourceSystemComponent() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSystemComponentRead,
		Schema:      componentSchema,
	}
}

func dataSourceSystemComponents() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSystemComponentsRead,
		Schema: map[string]*schema.Schema{
			"components": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: componentSchema,
				},
			},
		},
	}
}

func dataSourceSystemComponentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	component, err := config.client.GetSystemComponent(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(component.ID))
	if err := d.Set("name", component.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", component.Type); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("stage", component.Stage); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("cloud_providers", component.Providers); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("os_types", component.OsTypes); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("content", flattenComponentContent(component.Content)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func dataSourceSystemComponentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	res, err := config.client.GetSystemComponents()
	if err != nil {
		return diag.FromErr(err)
	}

	components := make([]map[string]interface{}, 0)
	for v := range res {
		components = append(components, map[string]interface{}{
			"id":              res[v].ID,
			"name":            res[v].Name,
			"type":            res[v].Type,
			"stage":           res[v].Stage,
			"cloud_providers": *res[v].Providers,
			"os_types":        *res[v].OsTypes,
			"content":         flattenComponentContent(res[v].Content),
		})
	}

	if err := d.Set("components", components); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) // nolint: gomnd

	return diags
}
