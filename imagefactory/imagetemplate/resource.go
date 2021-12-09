// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/sdk"
)

func Resource() *schema.Resource {
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

	config := m.(*config.Config)

	input := sdk.NewTemplate{
		Name:           graphql.String(d.Get("name").(string)),
		DistributionId: graphql.String(d.Get("distribution_id").(string)),
		Provider:       graphql.Provider(d.Get("cloud_provider").(string)),
		Config:         *expandTemplateConfig(d.Get("config").([]interface{})),
	}
	if len(d.Get("description").(string)) > 0 {
		description := graphql.String(d.Get("description").(string))
		input.Description = &description
	}
	template, err := config.APIClient.CreateTemplate(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(template.ID))

	resourceTemplateRead(ctx, d, m)

	return diags
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	var diags diag.Diagnostics

	config := m.(*config.Config)

	templateID := d.Id()

	template, err := config.APIClient.GetTemplate(templateID)
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

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	var diags diag.Diagnostics

	config := m.(*config.Config)

	templateID := d.Id()

	name := graphql.String(d.Get("name").(string))
	input := sdk.TemplateChanges{
		ID:     graphql.String(templateID),
		Name:   &name,
		Config: expandTemplateConfig(d.Get("config").([]interface{})),
	}
	if len(d.Get("description").(string)) > 0 {
		description := graphql.String(d.Get("description").(string))
		input.Description = &description
	}
	if _, err := config.APIClient.UpdateTemplate(input); err != nil {
		return diag.FromErr(err)
	}

	resourceTemplateRead(ctx, d, m)

	return diags
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*config.Config)

	templateID := d.Id()

	if err := config.APIClient.DeleteTemplate(templateID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
