// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var awsChinaAccountAccessResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"AWS_ACCESS_KEY_ID": {
			Type:     schema.TypeString,
			Required: true,
		},
		"AWS_SECRET_ACCESS_KEY": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var awsChinaAccountSchema = map[string]*schema.Schema{
	"alias": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"account_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"access": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     awsChinaAccountAccessResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     accountStateSchema,
	},
}

func ResourceAWSChina() *schema.Resource { // nolint: dupl
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return accountCreate(d, m, graphql.ProviderAWS, graphql.ScopeCHINA)
		},
		ReadContext:   resourceAccountRead,
		UpdateContext: resourceAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema:        awsChinaAccountSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
