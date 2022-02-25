// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package sdk

type API interface {
	// Account CRUD
	GetAccount(accountID string) (Account, error)
	CreateAccount(input NewAccount) (Account, error)
	UpdateAccount(input AccountChanges) (Account, error)
	DeleteAccount(accountID string) error

	// Component CRUD
	GetComponent(componentID string) (Component, error)
	CreateComponent(input NewComponent) (Component, error)
	UpdateComponent(input ComponentChanges) (Component, error)
	DeleteComponent(componentID string) error

	// ComponentVersion C--D
	CreateComponentVersion(input NewComponentContent) (Component, error)
	DeleteComponentVersion(componentID, version string) error

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
	GetSystemComponent(name string) (Component, error)
	GetCustomComponent(name, cloudProvider, stage string) (Component, error)

	// Variables CRUD
	GetVariable(name string) (Variable, error)
	CreateVariable(input NewVariable) (Variable, error)
	UpdateVariable(input NewVariable) (Variable, error)
	DeleteVariable(name string) error
}
