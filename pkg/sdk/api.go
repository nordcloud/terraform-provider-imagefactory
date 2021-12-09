// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package sdk

type API interface {
	// Account CRUD
	GetAccount(accountID string) (Account, error)
	CreateAccount(input NewAccount) (Account, error)
	UpdateAccount(input AccountChanges) (Account, error)
	DeleteAccount(accountID string) error

	// Distribution GET
	GetDistribution(name, cloudProvider string) (Distribution, error)

	// Template CRUD
	GetTemplate(templateID string) (Template, error)
	CreateTemplate(input NewTemplate) (Template, error)
	UpdateTemplate(input TemplateChanges) (Template, error)
	DeleteTemplate(templateID string) error

	// RoleBinding CRUD
	GetRoleBinding(roleBindingID string) (RoleBinding, error)
	CreateRoleBinding(input NewRoleBinding) (RoleBinding, error)
	UpdateRoleBinding(input RoleBindingChanges) (RoleBinding, error)
	DeleteRoleBinding(roleBindingID string) error

	// ApiKey GET
	GetApiKey(name string) (ApiKey, error)

	// Component GET
	GetSystemComponent(name, cloudProvider, stage string) (Component, error)
	GetCustomComponent(name, cloudProvider, stage string) (Component, error)
}
