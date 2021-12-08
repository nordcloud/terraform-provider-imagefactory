// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func expandAwsAccountAccess(in []interface{}) *graphql.AccountCredentials {
	accountCredentials := &graphql.AccountCredentials{}

	if len(in) == 0 {
		return accountCredentials
	}

	awsAccess := in[0].(map[string]interface{})
	roleExternalID := graphql.String(awsAccess["role_external_id"].(string))
	accountCredentials.Aws = &graphql.AWSCredentials{
		Roles: &[]graphql.AWSCredentialsRole{
			{
				Arn:        graphql.String(awsAccess["role_arn"].(string)),
				ExternalId: &roleExternalID,
			},
		},
	}

	return accountCredentials
}

func expandAzureSubscriptionAccess(in []interface{}) *graphql.AccountCredentials {
	accountCredentials := &graphql.AccountCredentials{}

	if len(in) == 0 {
		return accountCredentials
	}

	azureAccess := in[0].(map[string]interface{})
	accountCredentials.Azure = &graphql.AzureCredentials{
		ResourceGroupName:  graphql.String(azureAccess["resource_group_name"].(string)),
		TenantId:           graphql.String(azureAccess["tenant_id"].(string)),
		AppId:              graphql.String(azureAccess["app_id"].(string)),
		Password:           graphql.String(azureAccess["password"].(string)),
		StorageAccount:     graphql.String(azureAccess["storage_account"].(string)),
		StorageAccountKey:  graphql.String(azureAccess["storage_account_key"].(string)),
		SharedImageGallery: graphql.String(azureAccess["shared_image_gallery"].(string)),
	}

	return accountCredentials
}

func expandIMBCloudAccountAccess(in []interface{}) *graphql.AccountCredentials {
	accountCredentials := &graphql.AccountCredentials{}

	if len(in) == 0 {
		return accountCredentials
	}

	access := in[0].(map[string]interface{})
	accountCredentials.Ibmcloud = &graphql.IBMCloudCredentials{
		Apikey:            graphql.String(access["apikey"].(string)),
		Region:            graphql.String(access["region"].(string)),
		CosBucket:         graphql.String(access["cos_bucket"].(string)),
		ResourceGroupName: graphql.String(access["resource_group_name"].(string)),
		ResourceGroupId:   graphql.String(access["resource_group_id"].(string)),
	}

	return accountCredentials
}

func flattenAccountState(in *graphql.AccountState) map[string]string {
	out := map[string]string{
		"status": string(in.Status),
	}
	if in.Error != nil {
		out["error"] = string(*in.Error)
	}

	return out
}
