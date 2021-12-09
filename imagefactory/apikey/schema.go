package apikey

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var apiKeySchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
}
