// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package config

import (
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/sdk"
)

const (
	httpTimout    = time.Second * 30
	DefaultAPIURL = "https://api.imagefactory.nordcloudapp.com/graphql"
)

// Config is the configuration structure used to instantiate the Config provider.
type Config struct {
	APIClient sdk.API
}

func NewTerraformConfig(apiURL, apiKey string) *Config {
	httpClient := cleanhttp.DefaultClient()
	httpClient.Timeout = httpTimout

	gqlApiClient := sdk.NewGraphQLClient(httpClient, map[string]string{
		"x-api-key": apiKey,
	})

	return &Config{
		APIClient: sdk.NewAPIClient(gqlApiClient, apiURL),
	}
}

func TerraformConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_key": {
			Type:        schema.TypeString,
			Description: "ImageFactory API key",
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("IMAGEFACTORY_API_KEY", nil),
		},
		"api_url": {
			Type:        schema.TypeString,
			Description: "ImageFactory API URL",
			Optional:    true,
			Default:     DefaultAPIURL,
		},
	}
}
