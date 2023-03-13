// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package role

var validRuleActions = []string{
	"VIEW",
	"CREATE",
	"UPDATE",
	"DELETE",
	"ANY",
}

var validResourceActions = []string{
	"ACCOUNT",
	"API_KEY",
	"AUDIT_LOG",
	"COMPONENT",
	"NOTIFICATION_GROUP",
	"ROLE",
	"ROLE_BINDING",
	"TEMPLATE",
	"VARIABLE",
	"ANY",
}
