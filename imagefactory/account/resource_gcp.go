// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var gcpProjectAccessResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"private_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"auth_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_email": {
			Type:     schema.TypeString,
			Required: true,
		},
		"token_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"auth_provider_x509_cert_url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_x509_cert_url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"project_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"private_key_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var gcpProjectSchema = map[string]*schema.Schema{
	"alias": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"project_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"access": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     gcpProjectAccessResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     accountStateSchema,
	},
}

func ResourceGCP() *schema.Resource { // nolint: dupl
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return accountCreate(ctx, d, m, graphql.ProviderGCP)
		},
		ReadContext:   resourceAccountRead,
		UpdateContext: resourceAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema:        gcpProjectSchema,
	}
}
