// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import "github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"

var validAzureRegions = []string{
	"australiacentral",
	"australiacentral2",
	"australiaeast",
	"australiasoutheast",
	"brazilsouth",
	"canadacentral",
	"canadaeast",
	"centralindia",
	"centralus",
	"eastasia",
	"eastus",
	"eastus2",
	"francecentral",
	"francesouth",
	"japaneast",
	"japanwest",
	"koreacentral",
	"koreasouth",
	"northcentralus",
	"northeurope",
	"southafricanorth",
	"southafricawest",
	"southcentralus",
	"southeastasia",
	"southindia",
	"uksouth",
	"ukwest",
	"westcentralus",
	"westeurope",
	"westindia",
	"westus",
	"westus2",
}

var validNotificationTypes = []string{
	string(graphql.NotificationTypePUBSUB),
	string(graphql.NotificationTypeSNS),
	string(graphql.NotificationTypeWEBHOOK),
}

var validCloudProviders = []string{
	string(graphql.ProviderAWS),
	string(graphql.ProviderAZURE),
	string(graphql.ProviderGCP),
	string(graphql.ProviderIBMCLOUD),
	string(graphql.ProviderVMWARE),
}

var validScopes = []string{
	string(graphql.ScopePUBLIC),
	string(graphql.ScopeCHINA),
}
