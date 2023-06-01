// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import "github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"

var validNotificationTypes = []string{
	string(graphql.NotificationTypePUBSUB),
	string(graphql.NotificationTypeSNS),
	string(graphql.NotificationTypeWEBHOOK),
}

var validCloudProviders = []string{
	string(graphql.ProviderAWS),
	string(graphql.ProviderAZURE),
	string(graphql.ProviderEXOSCALE),
	string(graphql.ProviderGCP),
	string(graphql.ProviderIBMCLOUD),
	string(graphql.ProviderVMWARE),
}

var validScopes = []string{
	string(graphql.ScopePUBLIC),
	string(graphql.ScopeCHINA),
}
