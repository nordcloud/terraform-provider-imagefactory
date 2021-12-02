package imagefactory

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDistributions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDistributionsRead,
		Schema: map[string]*schema.Schema{
			"distributions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"provider": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDistributionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*Config)

	res := config.client.GetDistributions()

	var diags diag.Diagnostics

	distributions := make([]map[string]interface{}, 0)
	for _, d := range *res.Distributions.Results {
		a := make(map[string]interface{})
		a["id"] = d.ID
		a["name"] = d.Name
		a["provider"] = d.Provider
		distributions = append(distributions, a)
	}

	if err := d.Set("distributions", distributions); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) // nolint: gomnd

	return diags
}
