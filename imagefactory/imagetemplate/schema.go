// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

var additionalEbsVolumesResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"size": {
			Type:         schema.TypeInt,
			Required:     true,
			Description:  "EBS volume size between 1 and 10 GB.",
			ValidateFunc: validation.IntBetween(1, 10), // nolint: gomnd
		},
		"device_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Device name for the EBS volume. Available names for Linux are `/dev/sd[b-z]`, for Windows `xvd[b-z]`",
			ValidateFunc: validation.StringMatch(
				regexp.MustCompile("^(/dev/sd[b-z]|xvd[b-z])$"),
				"Must be a valid device name. For Linux it should be /dev/sd[b-z], for Windows xvd[b-z]",
			),
		},
	},
}

var awsTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"region": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"custom_image_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"additional_ebs_volumes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     additionalEbsVolumesResource,
			MaxItems: 10,
		},
		"kms_key_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The ID of the AWS KMS key that is used to encrypt the destination snapshot of the copied image. " +
				"To allow use of this key, onboarded master role `ImageFactoryMasterRole` must have permission to use the key. " +
				"You can use key ID, key ARN, alias name, or alias ARN.",
		},
	},
}

func validateVMImageDefinitionParameter(min, max int) schema.SchemaValidateDiagFunc {
	return func(val interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics

		v, ok := val.(string)
		if !ok {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "Invalid value type",
				Detail:        "Field value must be of type string",
				AttributePath: path,
			})
		}

		if len(v) < min || len(v) > max {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "Allowed values",
				Detail:        fmt.Sprintf("Expected length of the value to be in the range (%d - %d), got %s", min, max, v),
				AttributePath: path,
			})
		}

		if ok := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*[a-zA-Z0-9]$`).MatchString(v); !ok {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid value",
				Detail: "The value must contain only English letters, numbers, underscores and hyphens. " +
					"The value cannot begin or end with underscores or hyphens.",
				AttributePath: path,
			})
		}

		return diags
	}
}

var vmImageDefinitionAzureTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateVMImageDefinitionParameter(2, 80), // nolint: gomnd
		},
		"offer": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateVMImageDefinitionParameter(2, 64), // nolint: gomnd
		},
		"sku": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validateVMImageDefinitionParameter(2, 64), // nolint: gomnd
		},
	},
}

var azureTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"exclude_from_latest": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"eol_date_option": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Default value is set to true",
		},
		"replica_regions": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vm_image_definition": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     vmImageDefinitionAzureTemplateConfigResource,
		},
	},
}

var exoscaleTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"zone": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var templateComponentResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var templateNotificationsResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(validNotificationTypes, false),
		},
		"uri": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var templateTagsResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var templateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"cloud_account_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"aws": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     awsTemplateConfigResource,
		},
		"azure": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     azureTemplateConfigResource,
		},
		"exoscale": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     exoscaleTemplateConfigResource,
		},
		"build_components": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateComponentResource,
		},
		"test_components": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateComponentResource,
		},
		"notifications": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateNotificationsResource,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     templateTagsResource,
		},
		"scope": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(validScopes, false),
			Default:      graphql.ScopePUBLIC,
		},
		"disable_cyclical_rebuilds": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Disable cyclical rebuilds. " +
				"Cyclical rebuilds are rebuilds that are triggered automatically by ImageFactory when the source image is updated or " +
				"when there are security updates available for the packages installed in the image. If cyclical rebuilds are disabled, " +
				"the template will not be rebuilt automatically and the user will have to trigger the rebuild manually. " +
				"Default value is set to false.",
		},
	},
}

var templateStateSchema = &schema.Schema{
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

var templateSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"distribution_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"cloud_provider": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validCloudProviders, false),
	},
	"config": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     templateConfigResource,
	},
	"state": {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     templateStateSchema,
	},
}
