// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package apikey

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/sdk"
)

var apiKeySchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"secret": {
		Type:     schema.TypeString,
		Computed: true,
		Description: "The secret value will only be returned when creating the API key. " +
			"Please save this value because it won't be possible to get it later. " +
			"If you lost apiKey secret you have to create new ApiKey. " +
			"apikey can be used to access ImageFactory API by providing the `x-api-key` header in format: " +
			"`{API_KEY_ID}/{API_KEY_SECRET}` " +
			"apiKey does not grant any permissions to access API itself. " +
			"You have to create the `imagefactory_role_binding` and assign the access role to it to make it working.",
	},
}

func setProps(d *schema.ResourceData, key sdk.APIKey) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(key.ID))

	if err := d.Set("name", key.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret", key.Secret); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
