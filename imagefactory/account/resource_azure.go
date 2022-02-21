// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var azureSubscriptionAccessResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"resource_group_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"app_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"password": {
			Type:     schema.TypeString,
			Required: true,
		},
		"storage_account": {
			Type:     schema.TypeString,
			Required: true,
		},
		"storage_account_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"shared_image_gallery": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var azureSubscriptionSchema = map[string]*schema.Schema{
	"alias": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"subscription_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"access": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     azureSubscriptionAccessResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     accountStateSchema,
	},
}

func ResourceAzure() *schema.Resource { // nolint: dupl
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return accountCreate(d, m, graphql.ProviderAZURE, graphql.ScopePUBLIC)
		},
		ReadContext:   resourceAccountRead,
		UpdateContext: resourceAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema:        azureSubscriptionSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
