// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var distributionSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"cloud_provider": {
		Type:     schema.TypeString,
		Required: true,
	},
}

func dataSourceDistribution() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDistributionRead,
		Schema:      distributionSchema,
	}
}

func dataSourceDistributions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDistributionsRead,
		Schema: map[string]*schema.Schema{
			"distributions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: distributionSchema,
				},
			},
		},
	}
}

func dataSourceDistributionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	res, err := config.client.GetDistribution(d.Get("name").(string), d.Get("cloud_provider").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	distro := *res.Distributions.Results
	d.SetId(distro[0].ID)
	if err := d.Set("name", distro[0].Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_provider", distro[0].Provider); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func dataSourceDistributionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	res, err := config.client.GetDistributions()
	if err != nil {
		return diag.FromErr(err)
	}

	distributions := make([]map[string]interface{}, 0)
	for _, distro := range *res.Distributions.Results {
		distributions = append(distributions, map[string]interface{}{
			"id":             distro.ID,
			"name":           distro.Name,
			"cloud_provider": distro.Provider,
		})
	}

	if err := d.Set("distributions", distributions); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) // nolint: gomnd

	return diags
}
