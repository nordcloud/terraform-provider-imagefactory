// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_role" "admin" {
  name = "Template Admin Role"
  description = "A role to administrate templates (create/update/delete)"
  rules {
    resources = ["TEMPLATE", "COMPONENT", "VARIABLE"]
    actions   = ["ANY"]
  }
  rules {
    resources = ["ACCOUNT"]
    actions   = ["VIEW"]
  }
}

output "admin_role" {
  value = imagefactory_role.admin
}
