// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

var validCloudProviders = []string{
	"AWS",
	"AZURE",
	"GCP",
	"IBMCLOUD",
	"VMWARE",
}

var validStages = []string{
	"BUILD",
	"TEST",
}

var validOSTypes = []string{
	"LINUX",
	"WINDOWS",
}

var validProvisioners = []string{
	"POWERSHELL",
	"SHELL",
}
