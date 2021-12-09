// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

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

func (c APIClient) GetApiKey(name string) (ApiKey, error) {
	req, err := graphql.NewGetApiKeysRequest(c.apiURL, &graphql.GetApiKeysVariables{
		Input: graphql.CustomerApiKeysInput{},
	})
	if err != nil {
		return ApiKey{}, fmt.Errorf("getting api_key request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlAPI.Execute(req.Request, r); err != nil {
		return ApiKey{}, fmt.Errorf("getting api_key %w", err)
	}

	if r.ApiKeys.Results != nil {
		for _, k := range *r.ApiKeys.Results {
			if string(k.Name) == name {
				return ApiKey(k), nil
			}
		}
	}

	return ApiKey{}, fmt.Errorf("api_key '%s' not found", name)
}
