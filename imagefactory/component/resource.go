// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

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
		CreateContext: resourceComponentCreate,
		ReadContext:   resourceComponentRead,
		UpdateContext: resourceComponentUpdate,
		DeleteContext: resourceComponentDelete,
		Schema:        componentSchema,
	}
}

func resourceComponentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	input := sdk.NewComponent{
		Name:      graphql.String(d.Get("name").(string)),
		Stage:     graphql.ComponentStage(d.Get("stage").(string)),
		OsTypes:   expandOSTypes(d.Get("os_types").([]interface{})),
		Providers: expandProviders(d.Get("cloud_providers").([]interface{})),
		Content:   expandContent(d.Get("content").([]interface{})),
	}
	if len(d.Get("description").(string)) > 0 {
		description := graphql.String(d.Get("description").(string))
		input.Description = &description
	}
	component, err := c.APIClient.CreateComponent(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(component.ID))

	resourceComponentRead(ctx, d, m)

	return diags
}

func resourceComponentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	var diags diag.Diagnostics

	c := m.(*config.Config)

	componentID := d.Id()

	component, err := c.APIClient.GetComponent(componentID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", component.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", component.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("stage", component.Stage); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("os_types", flattenOSTypes(component.OsTypes)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_providers", flattenProviders(component.Providers)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(componentID)

	return diags
}

func resourceComponentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	componentID := d.Id()

	if d.HasChanges("name", "os_types", "cloud_providers", "description") {
		name := graphql.String(d.Get("name").(string))
		input := sdk.ComponentChanges{
			ID:        graphql.String(componentID),
			Name:      &name,
			OsTypes:   expandOSTypes(d.Get("os_types").([]interface{})),
			Providers: expandProviders(d.Get("cloud_providers").([]interface{})),
		}
		if len(d.Get("description").(string)) > 0 {
			description := graphql.String(d.Get("description").(string))
			input.Description = &description
		}
		if _, err := c.APIClient.UpdateComponent(input); err != nil {
			return diag.FromErr(err)
		}
	}

	// add new component version
	if d.HasChange("content") {
		in := d.Get("content").([]interface{})
		m := in[0].(map[string]interface{})

		active := graphql.Boolean(true)
		input := sdk.NewComponentContent{
			ID:                graphql.String(componentID),
			Active:            &active,
			Script:            graphql.String(m["script"].(string)),
			ScriptProvisioner: graphql.ShellScriptProvisioner(m["provisioner"].(string)),
		}

		if _, err := c.APIClient.CreateComponentVersion(input); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceComponentRead(ctx, d, m)
}

func resourceComponentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	componentID := d.Id()

	if err := c.APIClient.DeleteComponent(componentID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}