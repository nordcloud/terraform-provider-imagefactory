// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

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
