package imagefactory

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/httpclient"
)

// Config is the configuration structure used to instantiate the AutoPatcher provider.
type Config struct {
	ApiUrl         string
	ApiKey         string
	RequestTimeout time.Duration
	client         *Client

	context          context.Context
	terraformVersion string
	userAgent        string
}

func (c *Config) LoadAndValidate() error {
	httpClient := cleanhttp.DefaultClient()
	httpClient.Timeout = c.synchronousTimeout()
	tfUserAgent := httpclient.TerraformUserAgent(c.terraformVersion)
	providerVersion := "terraform-provider-imagefactory/dev"
	userAgent := fmt.Sprintf("%s %s", tfUserAgent, providerVersion)

	client := NewClient(c.ApiUrl, c.ApiKey, httpClient)
	client.userAgent = userAgent

	c.client = client

	return nil
}

func (c *Config) synchronousTimeout() time.Duration {
	if c.RequestTimeout == 0 {
		return 30 * time.Second
	}

	return c.RequestTimeout
}
