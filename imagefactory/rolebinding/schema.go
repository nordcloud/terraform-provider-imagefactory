// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package rolebinding

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var roleBindingSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"kind": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validBindingKinds, false),
	},
	"role_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"subject": {
		Type:     schema.TypeString,
		Required: true,
	},
}
