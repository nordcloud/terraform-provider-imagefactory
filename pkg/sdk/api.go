// Copyright 2021-2024 Nordcloud Oy or its affiliates. All Rights Reserved.

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

	// Role CRUD
	GetRole(roleID string) (Role, error)
	GetRoleByName(name string) (Role, error)
	CreateRole(input NewRole) (Role, error)
	UpdateRole(input RoleChanges) (Role, error)
	DeleteRole(roleID string) error

	// RoleBinding CRUD
	GetRoleBinding(roleBindingID string) (RoleBinding, error)
	CreateRoleBinding(input NewRoleBinding) (RoleBinding, error)
	UpdateRoleBinding(input RoleBindingChanges) (RoleBinding, error)
	DeleteRoleBinding(roleBindingID string) error

	// ApiKey CRUD
	GetAPIKey(apiKeyID string) (APIKey, error)
	GetAPIKeyByName(name string) (APIKey, error)
	CreateAPIKey(input NewAPIKey) (APIKey, error)
	UpdateAPIKey(input APIKeyChanges) (APIKey, error)
	DeleteAPIKey(apiKeyID string) error

	// Component GET
	GetSystemComponent(name string) (Component, error)
	GetCustomComponent(name, cloudProvider, stage string) (Component, error)

	// Variables CRUD
	GetVariable(name string) (Variable, error)
	CreateVariable(input NewVariable) (Variable, error)
	UpdateVariable(input NewVariable) (Variable, error)
	DeleteVariable(name string) error

	// Actions
	RebuildTemplatesUsingComponent(componentID string) error
}
