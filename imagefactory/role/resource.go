// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package role

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
		CreateContext: resourceRoleCreate,
		ReadContext:   resourceRoleRead,
		UpdateContext: resourceRoleUpdate,
		DeleteContext: resourceRoleDelete,
		Schema:        roleSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceRoleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*config.Config)

	input := sdk.NewRole{
		Name:  graphql.String(d.Get("name").(string)),
		Rules: expandRoleRules(d.Get("rules").([]interface{})),
	}
	if len(d.Get("description").(string)) > 0 {
		description := graphql.String(d.Get("description").(string))
		input.Description = &description
	}

	role, err := c.APIClient.CreateRole(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, role)
}

func resourceRoleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	role, err := c.APIClient.GetRole(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, role)
}

func resourceRoleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	var (
		role sdk.Role
		err  error
	)
	if d.HasChanges("name", "description", "rules") {
		name := graphql.String(d.Get("name").(string))
		input := sdk.RoleChanges{
			ID:    graphql.String(d.Id()),
			Name:  &name,
			Rules: expandRoleRules(d.Get("rules").([]interface{})),
		}
		if len(d.Get("description").(string)) > 0 {
			description := graphql.String(d.Get("description").(string))
			input.Description = &description
		}

		role, err = c.APIClient.UpdateRole(input)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return setProps(d, role)
}

func resourceRoleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	if err := c.APIClient.DeleteRole(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setProps(d *schema.ResourceData, rb sdk.Role) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(rb.ID))

	if err := d.Set("name", rb.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", rb.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rules", flattenRoleRules(rb.Rules)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
