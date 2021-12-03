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

func (c Client) GetSystemComponent(name, cloudProviders string) (*graphql.GetSystemComponentsResponse, error) {
	req, err := graphql.NewGetSystemComponentsRequest(c.endpoint, &graphql.GetSystemComponentsVariables{
		Input: graphql.ComponentsInput{
			Filters: &graphql.ComponentsFilters{
				Filters: &[]graphql.ComponentsFilter{
					{
						Field:  graphql.ComponentAttributeNAME,
						Values: &[]graphql.String{graphql.String(name)},
					},
					{
						Field:  graphql.ComponentAttributePROVIDERS,
						Values: &[]graphql.String{graphql.String(cloudProviders)},
					},
				},
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("getting component %w", err)
	}
	req.Header = http.Header{APIKeyHeader: []string{c.apiKey}}

	return req.Execute(c.httpClient)
}

func (c Client)  GetSystemComponents() (*graphql.GetSystemComponentsResponse, error) {
	req, err := graphql.NewGetSystemComponentsRequest(c.endpoint, &graphql.GetSystemComponentsVariables{})
	if err != nil {
		return nil, fmt.Errorf("getting system components %w", err)
	}
	req.Header = http.Header{APIKeyHeader: []string{c.apiKey}}

	return req.Execute(c.httpClient)
}