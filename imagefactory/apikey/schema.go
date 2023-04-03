// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package apikey

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/sdk"
)

const expiresAtDateFormat = "2006-01-02"

var apiKeySchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"expires_at": {
		Type:         schema.TypeString,
		Optional:     true,
		Description:  "API key expiration date in format: 2023-11-04",
		ValidateFunc: validateExpiresAtDateFormat(),
	},
	"secret": {
		Type:     schema.TypeString,
		Computed: true,
		Description: "The secret value will only be returned when creating the API key. " +
			"Please save this value because it won't be possible to get it later. " +
			"If you lost apiKey secret you have to create new ApiKey. " +
			"apikey can be used to access ImageFactory API by providing the `x-api-key` header in format: " +
			"`{API_KEY_ID}/{API_KEY_SECRET}` " +
			"apiKey does not grant any permissions to access API itself. " +
			"You have to create the `imagefactory_role_binding` and assign the access role to it to make it working.",
	},
}

func setProps(d *schema.ResourceData, key sdk.APIKey) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(string(key.ID))

	if err := d.Set("name", key.Name); err != nil {
		return diag.FromErr(err)
	}
	keySecret := key.Secret
	if keySecret == "" {
		keySecret = "***"
	}
	if err := d.Set("secret", keySecret); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("expires_at", formatExpiresAtDate((*string)(key.ExpiresAt))); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func formatExpiresAtDate(expiresAt *string) string {
	if expiresAt == nil {
		return ""
	}

	d, err := time.Parse(time.RFC3339, *expiresAt)
	if err != nil {
		return ""
	}

	return d.Format(expiresAtDateFormat)
}

func validateExpiresAtDateFormat() schema.SchemaValidateFunc { // nolint: staticcheck
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
			return warnings, errors
		}

		_, err := time.Parse(expiresAtDateFormat, v)
		if err != nil {
			message := "Invalid date format, an example of a valid date format: 2023-11-04"
			errors = append(errors, fmt.Errorf("invalid value for %s (%s)", k, message))
		}

		return warnings, errors
	}
}
