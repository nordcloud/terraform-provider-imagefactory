// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package apikey

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
		CreateContext: resourceAPIKeyCreate,
		ReadContext:   resourceAPIKeyRead,
		UpdateContext: schema.NoopContext,
		DeleteContext: resourceAPIKeyDelete,
		Schema:        apiKeySchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAPIKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*config.Config)

	input := sdk.NewAPIKey{
		Name: graphql.String(d.Get("name").(string)),
	}

	apiKey, err := c.APIClient.CreateAPIKey(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, apiKey)
}

func resourceAPIKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	apiKey, err := c.APIClient.GetAPIKey(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, apiKey)
}

func resourceAPIKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	if err := c.APIClient.DeleteAPIKey(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
