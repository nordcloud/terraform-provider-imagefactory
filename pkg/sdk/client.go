// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package sdk

import (
	"fmt"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

type APIClient struct {
	graphqlAPI GraphQLAPI
	apiURL     string
	userAgent  string
}

func NewAPIClient(gqlClient GraphQLAPI, apiURL string) API {
	return &APIClient{
		graphqlAPI: gqlClient,
		apiURL:     apiURL,
		userAgent:  "ImageFactorySDK",
	}
}

func (c APIClient) GetAccount(accountID string) (Account, error) { // nolint: dupl
	req, err := graphql.NewGetAccountRequest(c.apiURL, &graphql.GetAccountVariables{
		Input: graphql.CustomerAccountIdInput{
			AccountId: graphql.String(accountID),
		},
	})
	if err != nil {
		return Account{}, fmt.Errorf("getting account request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Account{}, fmt.Errorf("getting account %w", err)
	}

	return Account(r.Account), nil
}

func (c APIClient) CreateAccount(input NewAccount) (Account, error) {
	req, err := graphql.NewCreateAccountRequest(c.apiURL, &graphql.CreateAccountVariables{
		Input: graphql.NewAccount(input),
	})
	if err != nil {
		return Account{}, fmt.Errorf("getting create account request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Account{}, fmt.Errorf("creating account %w", err)
	}

	return Account(r.CreateAccount), nil
}

func (c APIClient) UpdateAccount(input AccountChanges) (Account, error) {
	req, err := graphql.NewUpdateAccountRequest(c.apiURL, &graphql.UpdateAccountVariables{
		Input: graphql.AccountChanges(input),
	})
	if err != nil {
		return Account{}, fmt.Errorf("getting update account request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Account{}, fmt.Errorf("updating account %w", err)
	}

	return Account(r.UpdateAccount), nil
}

func (c APIClient) DeleteAccount(accountID string) error {
	req, err := graphql.NewDeleteAccountRequest(c.apiURL, &graphql.DeleteAccountVariables{
		Input: graphql.CustomerAccountIdInput{
			AccountId: graphql.String(accountID),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete account request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting account %w", err)
	}

	return nil
}

func (c APIClient) GetComponent(componentID string) (Component, error) { // nolint: dupl
	req, err := graphql.NewGetComponentRequest(c.apiURL, &graphql.GetComponentVariables{
		Input: graphql.GetComponentInput{
			ComponentId: graphql.String(componentID),
		},
	})
	if err != nil {
		return Component{}, fmt.Errorf("getting component request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Component{}, fmt.Errorf("getting component %w", err)
	}

	return Component(r.Component), nil
}

func (c APIClient) CreateComponent(input NewComponent) (Component, error) {
	req, err := graphql.NewCreateComponentRequest(c.apiURL, &graphql.CreateComponentVariables{
		Input: graphql.NewComponent(input),
	})
	if err != nil {
		return Component{}, fmt.Errorf("getting create component request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Component{}, fmt.Errorf("creating component %w", err)
	}

	return Component(r.CreateComponent), nil
}

func (c APIClient) UpdateComponent(input ComponentChanges) (Component, error) {
	req, err := graphql.NewUpdateComponentRequest(c.apiURL, &graphql.UpdateComponentVariables{
		Input: graphql.ComponentChanges(input),
	})
	if err != nil {
		return Component{}, fmt.Errorf("getting update component request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Component{}, fmt.Errorf("updating component %w", err)
	}

	return Component(r.UpdateComponent), nil
}

func (c APIClient) DeleteComponent(componentID string) error {
	req, err := graphql.NewDeleteComponentRequest(c.apiURL, &graphql.DeleteComponentVariables{
		Input: graphql.ComponentIdInput{
			ComponentId: graphql.String(componentID),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete component request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting component %w", err)
	}

	return nil
}

func (c APIClient) CreateComponentVersion(input NewComponentContent) (Component, error) {
	req, err := graphql.NewCreateComponentVersionRequest(c.apiURL, &graphql.CreateComponentVersionVariables{
		Input: graphql.NewComponentContent(input),
	})
	if err != nil {
		return Component{}, fmt.Errorf("getting create component version request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Component{}, fmt.Errorf("creating component version %w", err)
	}

	return Component(r.CreateComponentVersion), nil
}

func (c APIClient) DeleteComponentVersion(componentID, version string) error {
	req, err := graphql.NewDeleteComponentVersionRequest(c.apiURL, &graphql.DeleteComponentVersionVariables{
		Input: graphql.ComponentVersionIdInput{
			ComponentId: graphql.String(componentID),
			Version:     graphql.String(version),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete component version request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting component version %w", err)
	}

	return nil
}

func (c APIClient) GetDistribution(name, cloudProvider string) (Distribution, error) {
	limit := graphql.Int(1)
	req, err := graphql.NewGetDistributionsRequest(c.apiURL, &graphql.GetDistributionsVariables{
		Input: graphql.DistributionsInput{
			Filters: &graphql.DistributionsFilters{
				Filters: &[]graphql.DistributionsFilter{
					{
						Field:  graphql.DistributionAttributeNAME,
						Values: &[]graphql.String{graphql.String(name)},
					},
					{
						Field:  graphql.DistributionAttributePROVIDER,
						Values: &[]graphql.String{graphql.String(cloudProvider)},
					},
				},
			},
			Limit: &limit,
		},
	})
	if err != nil {
		return Distribution{}, fmt.Errorf("getting distribution request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Distribution{}, fmt.Errorf("getting distribution %w", err)
	}

	if r.Distributions.Results == nil || len(*r.Distributions.Results) == 0 {
		return Distribution{}, fmt.Errorf("distribution '%s' in cloud provider '%s' not found", name, cloudProvider)
	}

	result := *r.Distributions.Results

	return Distribution(result[0]), nil
}

func (c APIClient) GetTemplate(templateID string) (Template, error) { // nolint: dupl
	req, err := graphql.NewGetTemplateRequest(c.apiURL, &graphql.GetTemplateVariables{
		Input: graphql.CustomerTemplateIdInput{
			TemplateId: graphql.String(templateID),
		},
	})
	if err != nil {
		return Template{}, fmt.Errorf("getting template request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Template{}, fmt.Errorf("getting template %w", err)
	}

	return Template(r.Template), nil
}

func (c APIClient) CreateTemplate(input NewTemplate) (Template, error) {
	req, err := graphql.NewCreateTemplateRequest(c.apiURL, &graphql.CreateTemplateVariables{
		Input: graphql.NewTemplate(input),
	})
	if err != nil {
		return Template{}, fmt.Errorf("getting create template request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Template{}, fmt.Errorf("creating template %w", err)
	}

	return Template(r.CreateTemplate), nil
}

func (c APIClient) UpdateTemplate(input TemplateChanges) (Template, error) {
	req, err := graphql.NewUpdateTemplateRequest(c.apiURL, &graphql.UpdateTemplateVariables{
		Input: graphql.TemplateChanges(input),
	})
	if err != nil {
		return Template{}, fmt.Errorf("getting update template request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Template{}, fmt.Errorf("updating template %w", err)
	}

	return Template(r.UpdateTemplate), nil
}

func (c APIClient) DeleteTemplate(templateID string) error {
	req, err := graphql.NewDeleteTemplateRequest(c.apiURL, &graphql.DeleteTemplateVariables{
		Input: graphql.CustomerTemplateIdInput{
			TemplateId: graphql.String(templateID),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete template request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting template %w", err)
	}

	return nil
}

func (c APIClient) GetRole(roleID string) (Role, error) { // nolint: dupl
	req, err := graphql.NewGetRoleRequest(c.apiURL, &graphql.GetRoleVariables{
		Input: graphql.CustomerRoleIdInput{
			RoleId: graphql.String(roleID),
		},
	})
	if err != nil {
		return Role{}, fmt.Errorf("getting role request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Role{}, fmt.Errorf("getting role %w", err)
	}

	return Role(r.Role), nil
}

func (c APIClient) GetRoleByName(name string) (Role, error) {
	limit := graphql.Int(1)
	req, err := graphql.NewGetRolesRequest(c.apiURL, &graphql.GetRolesVariables{
		Input: graphql.CustomerRolesInput{
			Filters: &graphql.CustomerRolesFilters{
				Filters: &[]graphql.CustomerRolesFilter{
					{
						Field:  graphql.CustomerRolesAttributeNAME,
						Values: &[]graphql.String{graphql.String(name)},
					},
				},
			},
			Limit: &limit,
		},
	})
	if err != nil {
		return Role{}, fmt.Errorf("getting role request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Role{}, fmt.Errorf("getting role %w", err)
	}

	if r.Roles.Results == nil || len(*r.Roles.Results) == 0 {
		return Role{}, fmt.Errorf("role '%s' not found", name)
	}

	result := *r.Roles.Results

	return Role(result[0]), nil
}

func (c APIClient) CreateRole(input NewRole) (Role, error) {
	req, err := graphql.NewCreateRoleRequest(c.apiURL, &graphql.CreateRoleVariables{
		Input: graphql.NewRole(input),
	})
	if err != nil {
		return Role{}, fmt.Errorf("getting create role request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Role{}, fmt.Errorf("creating role %w", err)
	}

	return Role(r.CreateRole), nil
}

func (c APIClient) UpdateRole(input RoleChanges) (Role, error) {
	req, err := graphql.NewUpdateRoleRequest(c.apiURL, &graphql.UpdateRoleVariables{
		Input: graphql.RoleChanges(input),
	})
	if err != nil {
		return Role{}, fmt.Errorf("getting update role request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Role{}, fmt.Errorf("updating role %w", err)
	}

	return Role(r.UpdateRole), nil
}

func (c APIClient) DeleteRole(roleID string) error {
	req, err := graphql.NewDeleteRoleRequest(c.apiURL, &graphql.DeleteRoleVariables{
		Input: graphql.CustomerRoleIdInput{
			RoleId: graphql.String(roleID),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete role request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting role %w", err)
	}

	return nil
}

func (c APIClient) GetRoleBinding(roleBindingID string) (RoleBinding, error) { // nolint: dupl
	req, err := graphql.NewGetRoleBindingRequest(c.apiURL, &graphql.GetRoleBindingVariables{
		Input: graphql.CustomerRoleBindingIdInput{
			RoleBindingId: graphql.String(roleBindingID),
		},
	})
	if err != nil {
		return RoleBinding{}, fmt.Errorf("getting roleBinding request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return RoleBinding{}, fmt.Errorf("getting roleBinding %w", err)
	}

	return RoleBinding(r.RoleBinding), nil
}

func (c APIClient) CreateRoleBinding(input NewRoleBinding) (RoleBinding, error) {
	req, err := graphql.NewCreateRoleBindingRequest(c.apiURL, &graphql.CreateRoleBindingVariables{
		Input: graphql.NewRoleBinding(input),
	})
	if err != nil {
		return RoleBinding{}, fmt.Errorf("getting create roleBinding request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return RoleBinding{}, fmt.Errorf("creating roleBinding %w", err)
	}

	return RoleBinding(r.CreateRoleBinding), nil
}

func (c APIClient) UpdateRoleBinding(input RoleBindingChanges) (RoleBinding, error) {
	req, err := graphql.NewUpdateRoleBindingRequest(c.apiURL, &graphql.UpdateRoleBindingVariables{
		Input: graphql.RoleBindingChanges(input),
	})
	if err != nil {
		return RoleBinding{}, fmt.Errorf("getting update roleBinding request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return RoleBinding{}, fmt.Errorf("updating roleBinding %w", err)
	}

	return RoleBinding(r.UpdateRoleBinding), nil
}

func (c APIClient) DeleteRoleBinding(roleBindingID string) error {
	req, err := graphql.NewDeleteRoleBindingRequest(c.apiURL, &graphql.DeleteRoleBindingVariables{
		Input: graphql.CustomerRoleBindingIdInput{
			RoleBindingId: graphql.String(roleBindingID),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete roleBinding request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting roleBinding %w", err)
	}

	return nil
}

func (c APIClient) GetAPIKey(apiKeyID string) (APIKey, error) { // nolint: dupl
	req, err := graphql.NewGetApiKeyRequest(c.apiURL, &graphql.GetApiKeyVariables{
		Input: graphql.CustomerApiKeyIdInput{
			ApiKeyId: graphql.String(apiKeyID),
		},
	})
	if err != nil {
		return APIKey{}, fmt.Errorf("getting role request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return APIKey{}, fmt.Errorf("getting role %w", err)
	}

	return APIKey(r.ApiKey), nil
}

func (c APIClient) GetAPIKeyByName(name string) (APIKey, error) {
	req, err := graphql.NewGetApiKeysRequest(c.apiURL, &graphql.GetApiKeysVariables{
		Input: graphql.CustomerApiKeysInput{},
	})
	if err != nil {
		return APIKey{}, fmt.Errorf("getting api_key request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return APIKey{}, fmt.Errorf("getting api_key %w", err)
	}

	if r.ApiKeys.Results != nil {
		for i := range *r.ApiKeys.Results {
			apiKey := (*r.ApiKeys.Results)[i]
			if string(apiKey.Name) == name {
				return APIKey(apiKey), nil
			}
		}
	}

	return APIKey{}, fmt.Errorf("api_key '%s' not found", name)
}

func (c APIClient) CreateAPIKey(input NewAPIKey) (APIKey, error) {
	req, err := graphql.NewCreateApiKeyRequest(c.apiURL, &graphql.CreateApiKeyVariables{
		Input: graphql.NewApiKey(input),
	})
	if err != nil {
		return APIKey{}, fmt.Errorf("getting create api key request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return APIKey{}, fmt.Errorf("creating api key %w", err)
	}

	return APIKey(r.CreateApiKey), nil
}

func (c APIClient) UpdateAPIKey(input APIKeyChanges) (APIKey, error) {
	req, err := graphql.NewUpdateApiKeyRequest(c.apiURL, &graphql.UpdateApiKeyVariables{
		Input: graphql.ApiKeyChanges(input),
	})
	if err != nil {
		return APIKey{}, fmt.Errorf("getting update api key request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return APIKey{}, fmt.Errorf("updating api key %w", err)
	}

	return APIKey(r.UpdateApiKey), nil
}

func (c APIClient) DeleteAPIKey(apiKeyID string) error {
	req, err := graphql.NewDeleteApiKeyRequest(c.apiURL, &graphql.DeleteApiKeyVariables{
		Input: graphql.CustomerApiKeyIdInput{
			ApiKeyId: graphql.String(apiKeyID),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete role request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting role %w", err)
	}

	return nil
}

func (c APIClient) GetSystemComponent(name string) (Component, error) {
	a := graphql.Boolean(true)
	req, err := graphql.NewGetComponentsRequest(c.apiURL, &graphql.GetComponentsVariables{
		Input: graphql.ComponentsInput{
			IncludeSystem: &a,
			Filters: &graphql.ComponentsFilters{
				Filters: &[]graphql.ComponentsFilter{
					{
						Field:  graphql.ComponentAttributeNAME,
						Values: &[]graphql.String{graphql.String(name)},
					},
					{
						Field:  graphql.ComponentAttributeTYPE,
						Values: &[]graphql.String{graphql.String("SYSTEM")},
					},
				},
			},
		},
	})
	if err != nil {
		return Component{}, fmt.Errorf("getting component request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Component{}, fmt.Errorf("getting component %w", err)
	}

	if r.Components.Results == nil || len(*r.Components.Results) == 0 {
		return Component{}, fmt.Errorf("component '%s' not found", name)
	}

	result := *r.Components.Results

	return Component(result[0]), nil
}

func (c APIClient) GetCustomComponent(name, cloudProvider, stage string) (Component, error) {
	req, err := graphql.NewGetComponentsRequest(c.apiURL, &graphql.GetComponentsVariables{
		Input: graphql.ComponentsInput{
			Filters: &graphql.ComponentsFilters{
				Filters: &[]graphql.ComponentsFilter{
					{
						Field:  graphql.ComponentAttributeNAME,
						Values: &[]graphql.String{graphql.String(name)},
					},
					{
						Field:  graphql.ComponentAttributePROVIDERS,
						Values: &[]graphql.String{graphql.String(cloudProvider)},
					},
					{
						Field:  graphql.ComponentAttributeSTAGE,
						Values: &[]graphql.String{graphql.String(stage)},
					},
				},
			},
		},
	})
	if err != nil {
		return Component{}, fmt.Errorf("getting component request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Component{}, fmt.Errorf("getting component %w", err)
	}

	if r.Components.Results == nil || len(*r.Components.Results) == 0 {
		return Component{}, fmt.Errorf("component '%s' in cloud provider '%s' and stage '%s' not found", name, cloudProvider, stage)
	}

	result := *r.Components.Results

	return Component(result[0]), nil
}

func (c APIClient) GetVariable(name string) (Variable, error) {
	req, err := graphql.NewGetVariableRequest(c.apiURL, &graphql.GetVariableVariables{
		Input: graphql.CustomerVariableNameInput{VariableName: graphql.String(name)},
	})
	if err != nil {
		return Variable{}, fmt.Errorf("getting variable request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Variable{}, fmt.Errorf("getting variable %w", err)
	}

	return Variable(r.Variable), nil
}

func (c APIClient) CreateVariable(input NewVariable) (Variable, error) {
	req, err := graphql.NewCreateVariableRequest(c.apiURL, &graphql.CreateVariableVariables{
		Input: graphql.NewVariable(input),
	})
	if err != nil {
		return Variable{}, fmt.Errorf("getting create variable request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return Variable{}, fmt.Errorf("creating variable %w", err)
	}

	return Variable(r.CreateVariable), nil
}

func (c APIClient) UpdateVariable(input NewVariable) (Variable, error) {
	return c.CreateVariable(input)
}

func (c APIClient) DeleteVariable(name string) error {
	req, err := graphql.NewDeleteVariableRequest(c.apiURL, &graphql.DeleteVariableVariables{
		Input: graphql.CustomerVariableNameInput{VariableName: graphql.String(name)},
	})
	if err != nil {
		return fmt.Errorf("getting delete variable request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting variable %w", err)
	}

	return nil
}

func (c APIClient) RebuildTemplatesUsingComponent(componentID string) error {
	req, err := graphql.NewGetTemplatesRequest(c.apiURL, &graphql.GetTemplatesVariables{
		Input: graphql.CustomerTemplatesInput{
			Search: &graphql.Search{
				SearchValue: graphql.String(componentID),
				Algorithm:   graphql.SearchAlgorithmSUBSTRINGMATCH,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("getting templates request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return fmt.Errorf("getting templates %w", err)
	}
	if r.Templates.Results == nil {
		return nil
	}

	for i := range *r.Templates.Results {
		template := (*r.Templates.Results)[i]

		// ignore templates that have the component pinned to a specific version
		if isNotLatestComponent(componentID, template.Config) {
			continue
		}

		req, err := graphql.NewRebuildTemplateRequest(c.apiURL, &graphql.RebuildTemplateVariables{
			Input: graphql.CustomerTemplateIdInput{
				TemplateId: template.ID,
			},
		})
		if err != nil {
			return fmt.Errorf("getting rebuild template request %w", err)
		}

		r := &graphql.Mutation{}
		if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
			return fmt.Errorf("rebuilding template ID: %s err: %w", template.ID, err)
		}
	}

	return nil
}

func isNotLatestComponent(componentID string, config *graphql.TemplateConfig) bool {
	const latest = "latest"

	if config == nil {
		return false
	}

	for _, components := range []*[]graphql.TemplateComponent{config.BuildComponents, config.TestComponents} {
		if components != nil {
			for _, bc := range *components {
				if bc.ID == graphql.String(componentID) && bc.Version != nil && *bc.Version != latest {
					return true
				}
			}
		}
	}

	return false
}
