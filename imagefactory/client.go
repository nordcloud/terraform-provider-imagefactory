// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"fmt"
	"net/http"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

const (
	APIKeyHeader = "x-api-key"
)

type Client struct {
	httpClient      *http.Client
	graphqlExecutor *graphql.Executor
	endpoint        string
	apiKey          string
	userAgent       string
}

func NewClient(endpoint, apiKey string, httpClient *http.Client) *Client {
	c := &Client{
		httpClient:      httpClient,
		graphqlExecutor: graphql.NewExecutor(httpClient, endpoint, apiKey),
		endpoint:        endpoint,
		apiKey:          apiKey,
		userAgent:       "ImageFactorySDK",
	}

	return c
}

func (c Client) GetAccount(accountID string) (graphql.Account, error) { // nolint: dupl
	req, err := graphql.NewGetAccountRequest(c.endpoint, &graphql.GetAccountVariables{
		Input: graphql.CustomerAccountIdInput{
			AccountId: graphql.String(accountID),
		},
	})
	if err != nil {
		return graphql.Account{}, fmt.Errorf("getting account request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return graphql.Account{}, fmt.Errorf("getting account %w", err)
	}

	return r.Account, nil
}

func (c Client) CreateAccount(input graphql.NewAccount) (graphql.Account, error) {
	req, err := graphql.NewCreateAccountRequest(c.endpoint, &graphql.CreateAccountVariables{
		Input: input,
	})
	if err != nil {
		return graphql.Account{}, fmt.Errorf("getting create account request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return graphql.Account{}, fmt.Errorf("creating account %w", err)
	}

	return r.CreateAccount, nil
}

func (c Client) UpdateAccount(input graphql.AccountChanges) (graphql.Account, error) {
	req, err := graphql.NewUpdateAccountRequest(c.endpoint, &graphql.UpdateAccountVariables{
		Input: input,
	})
	if err != nil {
		return graphql.Account{}, fmt.Errorf("getting update account request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return graphql.Account{}, fmt.Errorf("updating account %w", err)
	}

	return r.UpdateAccount, nil
}

func (c Client) DeleteAccount(accountID string) error {
	req, err := graphql.NewDeleteAccountRequest(c.endpoint, &graphql.DeleteAccountVariables{
		Input: graphql.CustomerAccountIdInput{
			AccountId: graphql.String(accountID),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete account request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting account %w", err)
	}

	return nil
}

func (c Client) GetDistribution(name, cloudProvider string) (graphql.Distribution, error) {
	limit := graphql.Int(1)
	req, err := graphql.NewGetDistributionsRequest(c.endpoint, &graphql.GetDistributionsVariables{
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
		return graphql.Distribution{}, fmt.Errorf("getting distribution request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return graphql.Distribution{}, fmt.Errorf("getting distribution %w", err)
	}

	if r.Distributions.Results == nil || len(*r.Distributions.Results) == 0 {
		return graphql.Distribution{}, fmt.Errorf("distribution '%s' in cloud provider '%s' not found", name, cloudProvider)
	}

	result := *r.Distributions.Results

	return result[0], nil
}

func (c Client) GetDistributions() ([]graphql.Distribution, error) {
	req, err := graphql.NewGetDistributionsRequest(c.endpoint, &graphql.GetDistributionsVariables{})
	if err != nil {
		return nil, fmt.Errorf("getting distributions request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return nil, fmt.Errorf("getting distributions %w", err)
	}

	if r.Distributions.Results == nil || len(*r.Distributions.Results) == 0 {
		return []graphql.Distribution{}, nil
	}

	return *r.Distributions.Results, nil
}

func (c Client) GetTemplate(templateID string) (graphql.Template, error) { // nolint: dupl
	req, err := graphql.NewGetTemplateRequest(c.endpoint, &graphql.GetTemplateVariables{
		Input: graphql.CustomerTemplateIdInput{
			TemplateId: graphql.String(templateID),
		},
	})
	if err != nil {
		return graphql.Template{}, fmt.Errorf("getting template request %w", err)
	}

	r := &graphql.Query{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return graphql.Template{}, fmt.Errorf("getting template %w", err)
	}

	return r.Template, nil
}

func (c Client) CreateTemplate(input graphql.NewTemplate) (graphql.Template, error) {
	req, err := graphql.NewCreateTemplateRequest(c.endpoint, &graphql.CreateTemplateVariables{
		Input: input,
	})
	if err != nil {
		return graphql.Template{}, fmt.Errorf("getting create template request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return graphql.Template{}, fmt.Errorf("creating template %w", err)
	}

	return r.CreateTemplate, nil
}

func (c Client) UpdateTemplate(input graphql.TemplateChanges) (graphql.Template, error) {
	req, err := graphql.NewUpdateTemplateRequest(c.endpoint, &graphql.UpdateTemplateVariables{
		Input: input,
	})
	if err != nil {
		return graphql.Template{}, fmt.Errorf("getting update template request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return graphql.Template{}, fmt.Errorf("updating template %w", err)
	}

	return r.UpdateTemplate, nil
}

func (c Client) DeleteTemplate(templateID string) error {
	req, err := graphql.NewDeleteTemplateRequest(c.endpoint, &graphql.DeleteTemplateVariables{
		Input: graphql.CustomerTemplateIdInput{
			TemplateId: graphql.String(templateID),
		},
	})
	if err != nil {
		return fmt.Errorf("getting delete template request %w", err)
	}

	r := &graphql.Mutation{}
	if err := c.graphqlExecutor.Execute(req.Request, r); err != nil {
		return fmt.Errorf("deleting template %w", err)
	}

	return nil
}
