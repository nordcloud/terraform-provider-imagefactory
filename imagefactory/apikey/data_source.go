// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package apikey

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/sdk"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: read,
		Schema:      apiKeySchema,
	}
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	config := m.(*config.Config)

	apiKey, err := config.APIClient.GetApiKey(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, apiKey)
}

func setProps(d *schema.ResourceData, key sdk.ApiKey) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(key.ID))
	if err := d.Set("name", key.Name); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
