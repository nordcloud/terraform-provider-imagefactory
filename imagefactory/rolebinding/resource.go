// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

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
	c := m.(*config.Config)

	input := sdk.NewRoleBinding{
		Kind:    graphql.Kind(d.Get("kind").(string)),
		RoleId:  graphql.String(d.Get("role_id").(string)),
		Subject: graphql.String(d.Get("subject").(string)),
	}

	roleBinding, err := c.APIClient.CreateRoleBinding(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, roleBinding)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	roleBinding, err := c.APIClient.GetRoleBinding(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, roleBinding)
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	input := sdk.RoleBindingChanges{
		ID:     graphql.String(d.Id()),
		RoleId: graphql.String(d.Get("role_id").(string)),
	}

	roleBinding, err := c.APIClient.UpdateRoleBinding(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, roleBinding)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	if err := c.APIClient.DeleteRoleBinding(d.Id()); err != nil {
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
	if err := d.Set("role_id", rb.RoleId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("subject", rb.Subject); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
