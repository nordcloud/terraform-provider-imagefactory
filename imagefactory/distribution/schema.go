package distribution

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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

var distributionsSchema = map[string]*schema.Schema{
	"distributions": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: distributionSchema,
		},
	},
}
