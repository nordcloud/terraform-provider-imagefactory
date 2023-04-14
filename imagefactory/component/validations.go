// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package component

var validStages = []string{
	"BUILD",
	"TEST",
}

var validCloudProviders = []string{
	"AWS",
	"AZURE",
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
