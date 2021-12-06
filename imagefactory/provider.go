// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DefaultAPIURL = "https://api.imagefactory.nordcloudapp.com/graphql"

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
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
		},
		DataSourcesMap: map[string]*schema.Resource{
			"imagefactory_distributions": dataSourceDistributions(),
			"imagefactory_distribution":  dataSourceDistribution(),
			"imagefactory_system_components": dataSourceSystemComponents(),
			"imagefactory_system_component":  dataSourceSystemComponent(),
		},
		ResourcesMap:         map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	config := Config{
		APIURL: d.Get("api_url").(string),
		APIKey: d.Get("api_key").(string),
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, diag.FromErr(err)
	}

	return &config, diags
}
