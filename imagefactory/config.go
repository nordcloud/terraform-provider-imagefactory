package imagefactory

import (
	"time"

	"github.com/hashicorp/go-cleanhttp"
)

// Config is the configuration structure used to instantiate the AutoPatcher provider.
type Config struct {
	APIURL         string
	APIKey         string
	RequestTimeout time.Duration
	client         *Client
}

func (c *Config) LoadAndValidate() error {
	httpClient := cleanhttp.DefaultClient()
	httpClient.Timeout = c.synchronousTimeout()
	client := NewClient(c.APIURL, c.APIKey, httpClient)
	c.client = client

	return nil
}

func (c *Config) synchronousTimeout() time.Duration {
	if c.RequestTimeout == 0 {
		return 30 * time.Second // nolint: gomnd
	}

	return c.RequestTimeout
}
