---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "imagefactory_system_component Data Source - terraform-provider-imagefactory"
subcategory: ""
description: |-
  
---

# imagefactory_system_component (Data Source)



## Example Usage

```terraform
// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

data "imagefactory_system_component" "hardening_level_1" {
  name = "Hardening level 1"
}

output "system_component" {
  value = data.imagefactory_system_component.hardening_level_1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Read-Only

- `id` (String) The ID of this resource.
