// Copyright 2021-2025 Nordcloud Oy or its affiliates. All Rights Reserved.

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
	case graphql.ProviderEXOSCALE:
		cloudProviderKey = "organization_name"
	case graphql.ProviderGCP:
		cloudProviderKey = "project_id"
	default:
	}

	return cloudProviderKey
}

func accountCreate(d *schema.ResourceData, m interface{}, provider graphql.Provider, scope graphql.Scope) diag.Diagnostics {
	c := m.(*config.Config)

	alias := graphql.String(d.Get("alias").(string))
	input := sdk.NewAccount{
		Alias:           &alias,
		CloudProviderId: graphql.String(d.Get(getCloudProviderKeyName(provider)).(string)),
		Provider:        provider,
		Scope:           &scope,
	}

	switch provider {
	case graphql.ProviderAWS:
		input.Credentials = expandAwsAccountAccess(d.Get("access").([]interface{}), scope)
		properties, ok := d.GetOk("properties")
		if ok {
			input.Properties = expandAwsAccountProperties(properties.([]interface{}))
		}
	case graphql.ProviderAZURE:
		input.Credentials = expandAzureSubscriptionAccess(d.Get("access").([]interface{}))
	case graphql.ProviderEXOSCALE:
		input.Credentials = expandExoscaleOrganizationAccess(d.Get("access").([]interface{}))
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
	account, err := c.APIClient.CreateAccount(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, account)
}

func resourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	accountID := d.Id()

	account, err := c.APIClient.GetAccount(accountID)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, account)
}

func accountUpdate(d *schema.ResourceData, m interface{}, provider graphql.Provider, scope graphql.Scope) diag.Diagnostics { // nolint: dupl
	c := m.(*config.Config)

	accountID := d.Id()

	alias := graphql.String(d.Get("alias").(string))
	input := sdk.AccountChanges{
		ID:    graphql.String(accountID),
		Alias: &alias,
	}

	if d.HasChange("access") {
		var creds graphql.AccountCredentials

		switch provider {
		case graphql.ProviderAWS:
			creds = expandAwsAccountAccess(d.Get("access").([]interface{}), scope)
			properties, ok := d.GetOk("properties")
			if ok {
				input.Properties = expandAwsAccountProperties(properties.([]interface{}))
			}
		case graphql.ProviderAZURE:
			creds = expandAzureSubscriptionAccess(d.Get("access").([]interface{}))
		case graphql.ProviderEXOSCALE:
			creds = expandExoscaleOrganizationAccess(d.Get("access").([]interface{}))
		case graphql.ProviderGCP:
			creds = expandGcpOrganizationAccess(d.Get("access").([]interface{}))
		case graphql.ProviderIBMCLOUD:
			creds = expandIMBCloudAccountAccess(d.Get("access").([]interface{}))
		default:
		}

		input.Credentials = &creds
	}

	account, err := c.APIClient.UpdateAccount(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, account)
}

func resourceAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*config.Config)

	accountID := d.Id()

	if err := c.APIClient.DeleteAccount(accountID); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setProps(d *schema.ResourceData, a sdk.Account) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(a.ID))

	if err := d.Set("alias", a.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", a.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(getCloudProviderKeyName(a.Provider), a.CloudProviderId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", flattenAccountState(a.State)); err != nil {
		return diag.FromErr(err)
	}

	if a.Provider == graphql.ProviderAWS && a.Properties != nil {
		if err := d.Set("properties", []map[string]interface{}{
			flattenAccountProperties(a.Properties),
		}); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
