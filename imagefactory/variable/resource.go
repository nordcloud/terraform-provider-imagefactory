// Copyright 2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package variable

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
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: delete,
		Schema:        variableSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*config.Config)

	input := sdk.NewVariable{
		Name:  graphql.String(d.Get("name").(string)),
		Value: graphql.String(d.Get("value").(string)),
	}

	variable, err := config.APIClient.CreateVariable(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, variable)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	config := m.(*config.Config)

	roleBinding, err := config.APIClient.GetVariable(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, roleBinding)
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	config := m.(*config.Config)

	input := sdk.NewVariable{
		Name:  graphql.String(d.Get("name").(string)),
		Value: graphql.String(d.Get("value").(string)),
	}

	roleBinding, err := config.APIClient.UpdateVariable(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, roleBinding)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	config := m.(*config.Config)

	if err := config.APIClient.DeleteVariable(d.Get("name").(string)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setProps(d *schema.ResourceData, v sdk.Variable) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(v.Name))

	if err := d.Set("name", v.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("value", "***secret***"); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
