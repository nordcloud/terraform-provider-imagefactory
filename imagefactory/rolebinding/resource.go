// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package rolebinding

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
		Schema:        roleBindingSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*config.Config)

	input := sdk.NewRoleBinding{
		Kind:    graphql.Kind(d.Get("kind").(string)),
		Role:    graphql.Role(d.Get("role").(string)),
		Subject: graphql.String(d.Get("subject").(string)),
	}

	roleBinding, err := config.APIClient.CreateRoleBinding(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, roleBinding)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	config := m.(*config.Config)

	roleBinding, err := config.APIClient.GetRoleBinding(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, roleBinding)
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	config := m.(*config.Config)

	role := graphql.Role(d.Get("role").(string))
	input := sdk.RoleBindingChanges{
		ID:   graphql.String(d.Id()),
		Role: &role,
	}

	roleBinding, err := config.APIClient.UpdateRoleBinding(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, roleBinding)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	config := m.(*config.Config)

	if err := config.APIClient.DeleteRoleBinding(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setProps(d *schema.ResourceData, rb sdk.RoleBinding) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(rb.ID))
	if err := d.Set("kind", rb.Kind); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("role", rb.Role); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("subject", rb.Subject); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
