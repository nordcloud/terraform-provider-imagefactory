// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/sdk"
)

func getCloudProviderKeyName(provider graphql.Provider) string {
	var cloudProviderKey string
	switch provider {
	case graphql.ProviderAWS, graphql.ProviderIBMCLOUD:
		cloudProviderKey = "account_id"
	case graphql.ProviderAZURE:
		cloudProviderKey = "subscription_id"
	case graphql.ProviderGCP:
		cloudProviderKey = "project_id"
	default:
	}

	return cloudProviderKey
}

func accountCreate(ctx context.Context, d *schema.ResourceData, m interface{}, provider graphql.Provider) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*config.Config)

	alias := graphql.String(d.Get("alias").(string))
	input := sdk.NewAccount{
		Alias:           &alias,
		CloudProviderId: graphql.String(d.Get(getCloudProviderKeyName(provider)).(string)),
		Provider:        provider,
	}

	switch provider {
	case graphql.ProviderAWS:
		input.Credentials = expandAwsAccountAccess(d.Get("access").([]interface{}))
	case graphql.ProviderAZURE:
		input.Credentials = expandAzureSubscriptionAccess(d.Get("access").([]interface{}))
	case graphql.ProviderGCP:
		input.Credentials = expandGcpOrganizationAccess(d.Get("access").([]interface{}))
	case graphql.ProviderIBMCLOUD:
		input.Credentials = expandIMBCloudAccountAccess(d.Get("access").([]interface{}))
	default:
	}

	if len(d.Get("description").(string)) > 0 {
		description := graphql.String(d.Get("description").(string))
		input.Description = &description
	}
	account, err := config.APIClient.CreateAccount(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(string(account.ID))

	resourceAccountRead(ctx, d, m)

	return diags
}

func resourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	var diags diag.Diagnostics

	config := m.(*config.Config)

	accountID := d.Id()

	account, err := config.APIClient.GetAccount(accountID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("alias", account.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", account.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(getCloudProviderKeyName(account.Provider), account.CloudProviderId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", flattenAccountState(account.State)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(accountID)

	return diags
}

func resourceAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	var diags diag.Diagnostics

	config := m.(*config.Config)

	accountID := d.Id()

	alias := graphql.String(d.Get("alias").(string))
	input := sdk.AccountChanges{
		ID:    graphql.String(accountID),
		Alias: &alias,
	}
	if _, err := config.APIClient.UpdateAccount(input); err != nil {
		return diag.FromErr(err)
	}

	resourceAccountRead(ctx, d, m)

	return diags
}

func resourceAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*config.Config)

	accountID := d.Id()

	if err := config.APIClient.DeleteAccount(accountID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
