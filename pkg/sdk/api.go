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

	// ComponentVersion CREATE
	CreateComponentVersion(input NewComponentContent) (Component, error)

	// Distribution GET
	GetDistribution(name, cloudProvider string) (Distribution, error)
	GetDistributions() ([]Distribution, error)

	// Template CRUD
	GetTemplate(templateID string) (Template, error)
	CreateTemplate(input NewTemplate) (Template, error)
	UpdateTemplate(input TemplateChanges) (Template, error)
	DeleteTemplate(templateID string) error
}
