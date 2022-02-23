// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package account

import (
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func expandAwsAccountAccess(in []interface{}, scope graphql.Scope) graphql.AccountCredentials {
	accountCredentials := graphql.AccountCredentials{}

	if len(in) == 0 {
		return accountCredentials
	}

	access := in[0].(map[string]interface{})
	if scope == graphql.ScopePUBLIC {
		roleExternalID := graphql.String(access["role_external_id"].(string))
		accountCredentials.Aws = &graphql.AWSCredentials{
			Roles: &[]graphql.AWSCredentialsRole{
				{
					Arn:        graphql.String(access["role_arn"].(string)),
					ExternalId: &roleExternalID,
				},
			},
		}
	} else {
		accessKeyID := graphql.String(access["aws_access_key_id"].(string))
		secretAccessKey := graphql.String(access["aws_secret_access_key"].(string))
		accountCredentials.Aws = &graphql.AWSCredentials{
			Credentials: &graphql.AWSCredentialsAccessKey{
				AWSACCESSKEYID:     accessKeyID,
				AWSSECRETACCESSKEY: secretAccessKey,
			},
		}
	}

	return accountCredentials
}

func expandAzureSubscriptionAccess(in []interface{}) graphql.AccountCredentials {
	accountCredentials := graphql.AccountCredentials{}

	if len(in) == 0 {
		return accountCredentials
	}

	access := in[0].(map[string]interface{})
	accountCredentials.Azure = &graphql.AzureCredentials{
		ResourceGroupName:  graphql.String(access["resource_group_name"].(string)),
		TenantId:           graphql.String(access["tenant_id"].(string)),
		AppId:              graphql.String(access["app_id"].(string)),
		Password:           graphql.String(access["password"].(string)),
		StorageAccount:     graphql.String(access["storage_account"].(string)),
		StorageAccountKey:  graphql.String(access["storage_account_key"].(string)),
		SharedImageGallery: graphql.String(access["shared_image_gallery"].(string)),
	}

	return accountCredentials
}

func expandGcpOrganizationAccess(in []interface{}) graphql.AccountCredentials {
	accountCredentials := graphql.AccountCredentials{}

	if len(in) == 0 {
		return accountCredentials
	}

	access := in[0].(map[string]interface{})
	accountCredentials.Gcp = &graphql.GCPCredentials{
		Type:                    graphql.String(access["type"].(string)),
		PrivateKey:              graphql.String(access["private_key"].(string)),
		AuthUri:                 graphql.String(access["auth_uri"].(string)),
		ClientId:                graphql.String(access["client_id"].(string)),
		ClientEmail:             graphql.String(access["client_email"].(string)),
		TokenUri:                graphql.String(access["token_uri"].(string)),
		AuthProviderX509CertUrl: graphql.String(access["auth_provider_x509_cert_url"].(string)),
		ClientX509CertUrl:       graphql.String(access["client_x509_cert_url"].(string)),
		ProjectId:               graphql.String(access["project_id"].(string)),
		PrivateKeyId:            graphql.String(access["private_key_id"].(string)),
	}

	return accountCredentials
}

func expandIMBCloudAccountAccess(in []interface{}) graphql.AccountCredentials {
	accountCredentials := graphql.AccountCredentials{}

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

func expandAwsAccountProperties(in interface{}) *graphql.AccountCloudPropertiesInput {
	awsAccountProps := graphql.AccountCloudPropertiesInput{}

	if in == nil {
		return nil
	}

	props := in.([]interface{})[0].(map[string]interface{})
	s3Bucket := graphql.String(props["s3_bucket_name"].(string))
	region := graphql.String(props["region"].(string))

	awsAccountProps.AwsChinaRegionName = &region
	awsAccountProps.AwsChinaS3BucketName = &s3Bucket

	if props["aws_share_accounts"] != nil {
		var shareAccounts []graphql.String
		shareAccountsIn := props["aws_share_accounts"].([]interface{})
		for _, acc := range shareAccountsIn {
			shareAccounts = append(shareAccounts, graphql.String(acc.(string)))
		}

		awsAccountProps.AwsShareAccounts = &shareAccounts
	}

	return &awsAccountProps
}
