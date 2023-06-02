// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var exoscaleOrganizationAccessResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"api_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"api_secret": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var exoscaleOrganizationSchema = map[string]*schema.Schema{
	"alias": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"organization_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"access": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     exoscaleOrganizationAccessResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     accountStateSchema,
	},
}

func ResourceExoscale() *schema.Resource { // nolint: dupl
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return accountCreate(d, m, graphql.ProviderEXOSCALE, graphql.ScopePUBLIC)
		},
		ReadContext: resourceAccountRead,
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return accountUpdate(d, m, graphql.ProviderEXOSCALE, graphql.ScopePUBLIC)
		},
		DeleteContext: resourceAccountDelete,
		Schema:        exoscaleOrganizationSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
