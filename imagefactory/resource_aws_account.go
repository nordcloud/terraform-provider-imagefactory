// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

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
		Optional: true,
		Elem:     awsAccountAccessResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     accountStateSchema,
	},
}

func resourceAwsAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAwsAccountCreate,
		ReadContext:   resourceAwsAccountRead,
		UpdateContext: resourceAwsAccountUpdate,
		DeleteContext: resourceAwsAccountDelete,
		Schema:        awsAccountSchema,
	}
}

func resourceAwsAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	alias := graphql.String(d.Get("alias").(string))
	input := graphql.NewAccount{
		Alias:           &alias,
		CloudProviderId: graphql.String(d.Get("account_id").(string)),
		Provider:        graphql.ProviderAWS,
		Credentials:     expandAwsAccountAccess(d.Get("access").([]interface{})),
	}
	if len(d.Get("description").(string)) > 0 {
		description := graphql.String(d.Get("description").(string))
		input.Description = &description
	}
	account, err := config.client.CreateAccount(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(account.ID))

	resourceAwsAccountRead(ctx, d, m)

	return diags
}

func resourceAwsAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	accountID := d.Id()

	account, err := config.client.GetAccount(accountID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("alias", account.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", account.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", account.CloudProviderId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", flattenAccountState(account.State)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(accountID)

	return diags
}

func resourceAwsAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	var diags diag.Diagnostics

	config := m.(*Config)

	accountID := d.Id()

	alias := graphql.String(d.Get("alias").(string))
	input := graphql.AccountChanges{
		ID:    graphql.String(accountID),
		Alias: &alias,
	}
	if _, err := config.client.UpdateAccount(input); err != nil {
		return diag.FromErr(err)
	}

	resourceAwsAccountRead(ctx, d, m)

	return diags
}

func resourceAwsAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	accountID := d.Id()

	if err := config.client.DeleteAccount(accountID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
