// Copyright 2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package variable

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var variableSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"value": {
		Type:     schema.TypeString,
		Required: true,
	},
}
