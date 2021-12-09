// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package rolebinding

var validBindingKinds = []string{
	"USER",
	"API_KEY",
}

var validBindingRoles = []string{
	"READ_ONLY",
	"ADMIN",
	"SUPER_ADMIN",
}
