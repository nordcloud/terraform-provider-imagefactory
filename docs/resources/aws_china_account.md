---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "imagefactory_aws_china_account Resource - terraform-provider-imagefactory"
subcategory: ""
description: |-
  
---

# imagefactory_aws_china_account (Resource)



## Example Usage

```terraform
// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_aws_china_account" "aws_china_account" {
  alias       = "IF AWS China Account"
  description = "Account to distribute AWS images"
  account_id  = "123456789012"
  access {
    AWS_ACCESS_KEY_ID     = "A...B"
    AWS_SECRET_ACCESS_KEY = "A..SECRET..B"
  }
  properties {
    s3_bucket_name = "s3_bucket_name"
    region = "cn-north-1"
  }
}

output "aws_china_account" {
  value = imagefactory_aws_china_account.aws_account
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **access** (Block List, Min: 1) (see [below for nested schema](#nestedblock--access))
- **account_id** (String)
- **alias** (String)
- **properties** (Block List, Min: 1) (see [below for nested schema](#nestedblock--properties))

### Optional

- **description** (String)
- **id** (String) The ID of this resource.

### Read-Only

- **state** (Map of String)

<a id="nestedblock--access"></a>
### Nested Schema for `access`

Required:

- **aws_access_key_id** (String)
- **aws_secret_access_key** (String)


<a id="nestedblock--properties"></a>
### Nested Schema for `properties`

Required:

- **region** (String)
- **s3_bucket_name** (String)

## Import

Import is supported using the following syntax:

```shell
# Copyright 2022 Nordcloud Oy or its affiliates. All Rights Reserved.

terraform import imagefactory_aws_china_account.tf_name RESOURCE_ID
```