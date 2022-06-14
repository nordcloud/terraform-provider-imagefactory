// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var awsAccountAccessResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"role_arn": {
			Type:     schema.TypeString,
			Required: true,
		},
		"role_external_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var awsAccountProperties = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"aws_share_accounts": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"aws_share_organizations": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"aws_share_ous": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	},
}

var awsAccountSchema = map[string]*schema.Schema{
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
		Elem:     awsAccountAccessResource,
		Required: true,
	},
	"properties": {
		Type:     schema.TypeList,
		Required: true,
		Elem:     awsAccountProperties,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     accountStateSchema,
	},
}

func ResourceAWS() *schema.Resource { // nolint: dupl
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return accountCreate(d, m, graphql.ProviderAWS, graphql.ScopePUBLIC)
		},
		ReadContext:   resourceAccountRead,
		UpdateContext: resourceAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema:        awsAccountSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
