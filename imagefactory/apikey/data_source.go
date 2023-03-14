// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package apikey

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: read,
		Schema:      apiKeySchema,
	}
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*config.Config)

	apiKey, err := c.APIClient.GetAPIKeyByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, apiKey)
}
