package imagefactory

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DEFAULT_API_URL = "https://api.hbi.nordcloudapp.com/graphql"

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
				Default:     DEFAULT_API_URL,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"imagefactory_distributions": dataSourceDistributions(),
		},
		ResourcesMap: map[string]*schema.Resource{},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}

		return providerConfigure(d, provider, terraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, p *schema.Provider, terraformVersion string) (interface{}, error) {
	config := Config{
		terraformVersion: terraformVersion,
		ApiKey:           d.Get("api_key").(string),
		ApiUrl:           d.Get("api_url").(string),
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
