// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

package role

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/config"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: roleRead,
		Schema:      roleSchema,
	}
}

func roleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*config.Config)

	role, err := c.APIClient.GetRoleByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return setProps(d, role)
}
