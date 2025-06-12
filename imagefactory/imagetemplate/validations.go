// Copyright 2021-2025 Nordcloud Oy or its affiliates. All Rights Reserved.

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

var validEBSVolumeTypes = []string{
	string(graphql.EBSVolumeTypeGp2),
	string(graphql.EBSVolumeTypeGp3),
}
