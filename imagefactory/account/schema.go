// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var accountStateSchema = &schema.Schema{
	Type: schema.TypeString,
	Elem: map[string]*schema.Schema{
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"error": {
			Type:     schema.TypeString,
			Computed: true,
		},
	},
}
