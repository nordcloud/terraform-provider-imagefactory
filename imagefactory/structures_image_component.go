// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagefactory

import (
	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func flattenComponentContent(in *[]graphql.VersionedContent) []interface{} {
	if in == nil {
		return nil
	}
	var out = make([]interface{}, len(*in))
  for i, step := range *in {
		m := make(map[string]string)
		m["version"]= string(step.Version)
		out[i] = m
	}

	return out
}
