// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*config.Config)

	tplCfg, err := expandTemplateConfig(d.Get("config").([]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}

	input := sdk.NewTemplate{
		Name:           graphql.String(d.Get("name").(string)),
		DistributionId: graphql.String(d.Get("distribution_id").(string)),
		Provider:       graphql.Provider(d.Get("cloud_provider").(string)),
		Config:         *tplCfg,
	}
	if len(d.Get("description").(string)) > 0 {
		description := graphql.String(d.Get("description").(string))
		input.Description = &description
	}
	template, err := c.APIClient.CreateTemplate(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, template)
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	templateID := d.Id()

	template, err := c.APIClient.GetTemplate(templateID)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, template)
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	templateID := d.Id()
	name := graphql.String(d.Get("name").(string))

	tplCfg, err := expandTemplateConfig(d.Get("config").([]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}

	input := sdk.TemplateChanges{
		ID:     graphql.String(templateID),
		Name:   &name,
		Config: tplCfg,
	}
	if len(d.Get("description").(string)) > 0 {
		description := graphql.String(d.Get("description").(string))
		input.Description = &description
	}
	template, err := c.APIClient.UpdateTemplate(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, template)
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	templateID := d.Id()

	if err := c.APIClient.DeleteTemplate(templateID); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setProps(d *schema.ResourceData, t sdk.Template) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(t.ID))

	if err := d.Set("name", t.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", t.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_provider", t.Provider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("distribution_id", t.DistributionId); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", flattenTemplateState(t.State)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
