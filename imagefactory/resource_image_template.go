// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var componentResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var awsConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"region": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var configResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"test_components": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     componentResource,
		},
		"build_components": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     componentResource,
		},
		"aws": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     awsConfigResource,
		},
	},
}

var stateSchema = &schema.Schema{
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
		Elem:     configResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     stateSchema,
	},
}

func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateCreate,
		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateUpdate,
		DeleteContext: resourceTemplateDelete,
		Schema:        templateSchema,
	}
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	template, err := config.client.CreateTemplate(graphql.NewTemplate{
		Name:           graphql.String(d.Get("name").(string)),
		DistributionId: graphql.String(d.Get("distribution_id").(string)),
		Provider:       graphql.Provider(d.Get("cloud_provider").(string)),
		Config:         *expandTemplateConfig(d.Get("config").([]interface{})),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(template.ID))

	resourceTemplateRead(ctx, d, m)

	return diags
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	templateID := d.Id()

	template, err := config.client.GetTemplate(templateID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", template.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", template.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_provider", template.Provider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", flattenTemplateState(template.State)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(templateID)

	return diags
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceTemplateRead(ctx, d, m)
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	templateID := d.Id()

	if err := config.client.DeleteTemplate(templateID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
