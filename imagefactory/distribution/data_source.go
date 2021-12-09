// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package distribution

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: distributionRead,
		Schema:      distributionSchema,
	}
}

func distributionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*config.Config)

	distro, err := config.APIClient.GetDistribution(d.Get("name").(string), d.Get("cloud_provider").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(distro.ID))
	if err := d.Set("name", distro.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_provider", distro.Provider); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
