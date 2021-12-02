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
	"description": {
		Type:     schema.TypeString,
		Computed: true,
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
	config := m.(*Config)

	var diags diag.Diagnostics

	res := config.client.GetDistribution(d.Get("name").(string), d.Get("cloud_provider").(string))
	distro := *res.Distributions.Results
	d.Set("name", distro[0].Name)
	d.Set("cloud_provider", distro[0].Provider)
	d.SetId(distro[0].ID)

	return diags
}

func dataSourceDistributionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*Config)

	res := config.client.GetDistributions()

	var diags diag.Diagnostics

	distributions := make([]map[string]interface{}, 0)
	for _, distro := range *res.Distributions.Results {
		a := make(map[string]interface{})
		a["id"] = distro.ID
		a["name"] = distro.Name
		a["cloud_provider"] = distro.Provider
		distributions = append(distributions, a)
	}

	if err := d.Set("distributions", distributions); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) // nolint: gomnd

	return diags
}
