// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

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
	"PUB_SUB",
	"SNS",
	"WEB_HOOK",
}

var validCloudProviders = []string{
	"AWS",
	"AZURE",
	"GCP",
	"IBMCLOUD",
	"VMWARE",
}
