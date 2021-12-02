package imagefactory

import (
	"log"
	"net/http"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

type Client struct {
	httpClient *http.Client
	endpoint   string
	apiKey     string
	userAgent  string
}

func NewClient(endpoint string, apiKey string, httpClient *http.Client) *Client {
	c := &Client{
		httpClient: httpClient,
		endpoint:   endpoint,
		apiKey:     apiKey,
		userAgent:  "ImageFactorySDK",
	}

	return c
}

func (c Client) GetDistributions() *graphql.GetDistributionsResponse {
	req, err := graphql.NewGetDistributionsRequest(c.endpoint, &graphql.GetDistributionsVariables{})
	if err != nil {
		log.Fatal(err)
	}
	req.Header = http.Header{
		"x-api-key": []string{c.apiKey},
	}

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
