// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package distribution

import (
	"context"
	"strconv"
	"time"

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

func DataSources() *schema.Resource {
	return &schema.Resource{
		ReadContext: distributionsRead,
		Schema:      distributionsSchema,
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

func distributionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*config.Config)

	res, err := config.APIClient.GetDistributions()
	if err != nil {
		return diag.FromErr(err)
	}

	distributions := make([]map[string]interface{}, 0)
	for v := range res {
		distributions = append(distributions, map[string]interface{}{
			"id":             res[v].ID,
			"name":           res[v].Name,
			"cloud_provider": res[v].Provider,
		})
	}

	if err := d.Set("distributions", distributions); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) // nolint: gomnd

	return diags
}
