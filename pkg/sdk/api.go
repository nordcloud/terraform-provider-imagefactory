package sdk

type API interface {
	// Account CRUD
	GetAccount(accountID string) (Account, error)
	CreateAccount(input NewAccount) (Account, error)
	UpdateAccount(input AccountChanges) (Account, error)
	DeleteAccount(accountID string) error

	// Distribution GET
	GetDistribution(name, cloudProvider string) (Distribution, error)
	GetDistributions() ([]Distribution, error)

	// Template CRUD
	GetTemplate(templateID string) (Template, error)
	CreateTemplate(input NewTemplate) (Template, error)
	UpdateTemplate(input TemplateChanges) (Template, error)
	DeleteTemplate(templateID string) error
}
