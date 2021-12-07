// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var awsAccountCredentialsResource = &schema.Resource{
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

var accountCredentialsResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"aws": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     awsAccountCredentialsResource,
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

var accountSchema = map[string]*schema.Schema{
	"alias": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"cloud_provider_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"cloud_provider": {
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringInSlice([]string{
			"AWS",
			"AZURE",
			"GCP",
			"IBMCLOUD",
			"VMWARE",
		}, false),
	},
	"credentials": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     accountCredentialsResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     accountStateSchema,
	},
}

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAccountCreate,
		ReadContext:   resourceAccountRead,
		UpdateContext: resourceAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema:        accountSchema,
	}
}

func resourceAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	alias := graphql.String(d.Get("alias").(string))
	input := graphql.NewAccount{
		Alias:           &alias,
		CloudProviderId: graphql.String(d.Get("cloud_provider_id").(string)),
		Provider:        graphql.Provider(d.Get("cloud_provider").(string)),
		Credentials:     expandAccountCredentials(d.Get("credentials").([]interface{})),
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

	resourceAccountRead(ctx, d, m)

	return diags
}

func resourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	if err := d.Set("cloud_provider", account.Provider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_provider_id", account.CloudProviderId); err != nil {
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

	resourceAccountRead(ctx, d, m)

	return diags
}

func resourceAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := m.(*Config)

	accountID := d.Id()

	if err := config.client.DeleteAccount(accountID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
