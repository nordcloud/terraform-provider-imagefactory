// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Access struct {
	client *http.Client
	apiURL string
	apiKey string
}

func NewAccess(httpClient *http.Client, apiURL, apiKey string) *Access {
	return &Access{
		&http.Client{},
		apiURL,
		apiKey,
	}
}

func (g *Access) Execute(req *http.Request, result interface{}) error {
	req.Header.Add("x-api-key", g.apiKey)

	resp, err := execute(g.client, req)
	if err != nil {
		return err
	}

	if len(resp.Errors) != 0 {
		return fmt.Errorf("query to %s failed: %s", req.URL.Path, resp.Error())
	}

	fmt.Println(string(resp.Data))
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return fmt.Errorf("failed to unmarshal response JSON '%s': %s", resp.Data, err)
	}
	fmt.Println(result)

	return nil
}
