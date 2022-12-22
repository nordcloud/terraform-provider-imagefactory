// Copyright 2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package variable

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/helper/mutexkv"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/sdk"
)

// This is a global MutexKV for use within this plugin.
var variableMutexKV = mutexkv.NewMutexKV("variable")

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
	c := m.(*config.Config)
	variableName := d.Get("name").(string)
	variableValue := d.Get("value").(string)

	variableMutexKV.Lock(ctx, variableName)
	defer variableMutexKV.Unlock(ctx, variableName)

	input := sdk.NewVariable{
		Name:  graphql.String(variableName),
		Value: graphql.String(variableValue),
	}

	variable, err := c.APIClient.CreateVariable(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, variable)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)
	variableName := d.Get("name").(string)

	variable, err := c.APIClient.GetVariable(variableName)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, variable)
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)
	variableName := d.Get("name").(string)
	variableValue := d.Get("value").(string)

	variableMutexKV.Lock(ctx, variableName)
	defer variableMutexKV.Unlock(ctx, variableName)

	input := sdk.NewVariable{
		Name:  graphql.String(variableName),
		Value: graphql.String(variableValue),
	}

	variable, err := c.APIClient.UpdateVariable(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, variable)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*config.Config)
	variableName := d.Get("name").(string)

	variableMutexKV.Lock(ctx, variableName)
	defer variableMutexKV.Unlock(ctx, variableName)

	if err := c.APIClient.DeleteVariable(variableName); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setProps(d *schema.ResourceData, v sdk.Variable) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(v.Hash))

	if err := d.Set("name", v.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
