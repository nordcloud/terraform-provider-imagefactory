---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "imagefactory_imbcloud_account Resource - terraform-provider-imagefactory"
subcategory: ""
description: |-
  
---

# imagefactory_imbcloud_account (Resource)



## Example Usage

```terraform
// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_imbcloud_account" "imbcloud_account" {
  alias       = "IF IBMCloud Account"
  description = "IF IBMCloud Account to distribute IBMCloud images"
  account_id  = "1234567"
  access {
    apikey              = "APIKEY"
    region              = "eu-de"
    cos_bucket          = "nordcloudimagefactory-bucket"
    resource_group_name = "nordcloudimagefactory-rg"
    resource_group_id   = "asdfg12345zxcvb67890tyuiohjkllba"
  }
}

output "imbcloud_account" {
  value = imagefactory_imbcloud_account.imbcloud_account
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **account_id** (String)
- **alias** (String)

### Optional

- **access** (Block List) (see [below for nested schema](#nestedblock--access))
- **description** (String)
- **id** (String) The ID of this resource.

### Read-Only

- **state** (Map of String)

<a id="nestedblock--access"></a>
### Nested Schema for `access`

Required:

- **apikey** (String)
- **cos_bucket** (String)
- **region** (String)
- **resource_group_id** (String)
- **resource_group_name** (String)

## Import

Import is supported using the following syntax:

```shell
terraform import imagefactory_imbcloud_account.tf_name RESOURCE_ID
```