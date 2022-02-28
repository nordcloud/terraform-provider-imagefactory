// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

import (
	"context"
	"fmt"

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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceComponentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	return setProps(d, component)
}

func resourceComponentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	componentID := d.Id()

	component, err := c.APIClient.GetComponent(componentID)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, component)
}

func resourceComponentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	componentID := d.Id()

	var (
		err       error
		component sdk.Component
	)
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

		component, err = c.APIClient.UpdateComponent(input)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("content") {
		in := d.Get("content").([]interface{})
		m := in[0].(map[string]interface{})

		input := sdk.NewComponentContent{
			ID:                graphql.String(componentID),
			Script:            graphql.String(m["script"].(string)),
			ScriptProvisioner: graphql.ShellScriptProvisioner(m["provisioner"].(string)),
		}

		component, err = c.APIClient.CreateComponentVersion(input)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return setProps(d, component)
}

func resourceComponentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	componentID := d.Id()

	component, err := c.APIClient.GetComponent(componentID)
	if err != nil {
		return diag.FromErr(err)
	}

	if component.Content != nil && len(*component.Content) > 1 {
		for _, v := range *component.Content {
			if v.Latest {
				continue
			}

			if err := c.APIClient.DeleteComponentVersion(componentID, string(v.Version)); err != nil {
				return diag.FromErr(fmt.Errorf("deleting component version %s %w", v.Version, err))
			}
		}
	}

	if err := c.APIClient.DeleteComponent(componentID); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setProps(d *schema.ResourceData, c sdk.Component) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(c.ID))

	if err := d.Set("name", c.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", c.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("stage", c.Stage); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("os_types", flattenOSTypes(c.OsTypes)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_providers", flattenProviders(c.Providers)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
