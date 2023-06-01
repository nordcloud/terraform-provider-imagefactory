// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

var validStages = []string{
	"BUILD",
	"TEST",
}

var validCloudProviders = []string{
	"AWS",
	"AZURE",
	"EXOSCALE",
	"GCP",
	"IBMCLOUD",
	"VMWARE",
}

var validOSTypes = []string{
	"LINUX",
	"WINDOWS",
}

var validProvisioners = []string{
	"POWERSHELL",
	"SHELL",
	"ANSIBLE",
}
