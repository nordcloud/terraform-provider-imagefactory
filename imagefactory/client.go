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
	httpClient *http.Client
	endpoint   string
	apiKey     string
	userAgent  string
}

func NewClient(endpoint, apiKey string, httpClient *http.Client) *Client {
	c := &Client{
		httpClient: httpClient,
		endpoint:   endpoint,
		apiKey:     apiKey,
		userAgent:  "ImageFactorySDK",
	}

	return c
}

func (c Client) GetDistribution(name, cloudProvider string) (*graphql.GetDistributionsResponse, error) {
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
		},
	})
	if err != nil {
		return nil, fmt.Errorf("getting distribution %w", err)
	}
	req.Header = http.Header{APIKeyHeader: []string{c.apiKey}}

	return req.Execute(c.httpClient)
}

func (c Client) GetDistributions() (*graphql.GetDistributionsResponse, error) {
	req, err := graphql.NewGetDistributionsRequest(c.endpoint, &graphql.GetDistributionsVariables{})
	if err != nil {
		return nil, fmt.Errorf("getting distributions %w", err)
	}
	req.Header = http.Header{APIKeyHeader: []string{c.apiKey}}

	return req.Execute(c.httpClient)
}

func (c Client) GetTemplate(templateID string) (*graphql.GetTemplateResponse, error) {
	req, err := graphql.NewGetTemplateRequest(c.endpoint, &graphql.GetTemplateVariables{
		Input: graphql.CustomerTemplateIdInput{
			TemplateId: graphql.String(templateID),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("getting template %w", err)
	}
	req.Header = http.Header{APIKeyHeader: []string{c.apiKey}}

	return req.Execute(c.httpClient)
}

func (c Client) CreateTemplate(input graphql.NewTemplate) (*graphql.CreateTemplateResponse, error) {
	req, err := graphql.NewCreateTemplateRequest(c.endpoint, &graphql.CreateTemplateVariables{
		Input: input,
	})
	if err != nil {
		return nil, fmt.Errorf("creating template %w", err)
	}
	req.Header = http.Header{APIKeyHeader: []string{c.apiKey}}

	return req.Execute(c.httpClient)
}
