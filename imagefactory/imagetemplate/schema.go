// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

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
	},
}

func validateVMImageDefinitionParameter(min, max int) schema.SchemaValidateFunc { // nolint: staticcheck
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
			return warnings, errors
		}

		if len(v) < min || len(v) > max {
			errors = append(errors, fmt.Errorf("expected length of %s to be in the range (%d - %d), got %s", k, min, max, v))
		}

		if ok := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*[a-zA-Z0-9]$`).MatchString(v); !ok {
			message := "The value must contain only English letters, numbers, underscores and hyphens. " +
				"The value cannot begin or end with underscores or hyphens."
			errors = append(errors, fmt.Errorf("invalid value for %s (%s)", k, message))
		}

		return warnings, errors
	}
}

var vmImageDefinitionAzureTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validateVMImageDefinitionParameter(2, 80), // nolint: gomnd
		},
		"offer": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validateVMImageDefinitionParameter(2, 64), // nolint: gomnd
		},
		"sku": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validateVMImageDefinitionParameter(2, 64), // nolint: gomnd
		},
	},
}

var azureTemplateConfigResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"exclude_from_latest": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"replica_regions": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(validAzureRegions, false),
			},
		},
		"vm_image_definition": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     vmImageDefinitionAzureTemplateConfigResource,
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
