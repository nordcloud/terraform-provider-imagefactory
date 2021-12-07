// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func expandAccountCredentials(in []interface{}) *graphql.AccountCredentials {
	accountCredentials := &graphql.AccountCredentials{}

	if len(in) == 0 {
		return accountCredentials
	}

	m := in[0].(map[string]interface{})

	awsConfig := m["aws"].([]interface{})
	if len(awsConfig) > 0 {
		t := awsConfig[0].(map[string]interface{})
		roleExternalID := graphql.String(t["role_external_id"].(string))
		accountCredentials.Aws = &graphql.AWSCredentials{
			Roles: &[]graphql.AWSCredentialsRole{
				{
					Arn:        graphql.String(t["role_arn"].(string)),
					ExternalId: &roleExternalID,
				},
			},
		}
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
