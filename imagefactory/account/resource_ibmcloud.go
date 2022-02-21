// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var ibmCloudAccountAccessResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"apikey": {
			Type:     schema.TypeString,
			Required: true,
		},
		"region": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cos_bucket": {
			Type:     schema.TypeString,
			Required: true,
		},
		"resource_group_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"resource_group_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var ibmCloudAccountSchema = map[string]*schema.Schema{
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
		Elem:     ibmCloudAccountAccessResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     accountStateSchema,
	},
}

func ResourceIBMCloud() *schema.Resource { // nolint: dupl
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return accountCreate(d, m, graphql.ProviderIBMCLOUD, graphql.ScopePUBLIC)
		},
		ReadContext:   resourceAccountRead,
		UpdateContext: resourceAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema:        ibmCloudAccountSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
