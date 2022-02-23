// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/imagefactory/account"
	"github.com/nordcloud/terraform-provider-imagefactory/imagefactory/apikey"
	"github.com/nordcloud/terraform-provider-imagefactory/imagefactory/component"
	"github.com/nordcloud/terraform-provider-imagefactory/imagefactory/distribution"
	"github.com/nordcloud/terraform-provider-imagefactory/imagefactory/imagetemplate"
	"github.com/nordcloud/terraform-provider-imagefactory/imagefactory/rolebinding"
	"github.com/nordcloud/terraform-provider-imagefactory/imagefactory/variable"
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: config.TerraformConfigSchema(),
		DataSourcesMap: map[string]*schema.Resource{
			"imagefactory_distribution":     distribution.DataSource(),
			"imagefactory_api_key":          apikey.DataSource(),
			"imagefactory_system_component": component.DataSourceSystem(),
			"imagefactory_custom_component": component.DataSourceCustom(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"imagefactory_aws_account":        account.ResourceAWS(),
			"imagefactory_azure_subscription": account.ResourceAzure(),
			"imagefactory_gcp_project":        account.ResourceGCP(),
			"imagefactory_ibmcloud_account":   account.ResourceIBMCloud(),
			"imagefactory_aws_china_account":  account.ResourceAWSChina(),
			"imagefactory_custom_component":   component.Resource(),
			"imagefactory_template":           imagetemplate.Resource(),
			"imagefactory_role_binding":       rolebinding.Resource(),
			"imagefactory_variable":           variable.Resource(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	return config.NewTerraformConfig(
		d.Get("api_url").(string),
		d.Get("api_key").(string),
	), diags
}
