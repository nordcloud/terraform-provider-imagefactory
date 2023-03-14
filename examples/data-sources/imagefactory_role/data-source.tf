// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

data "imagefactory_role" "admin_role" {
  name           = "Admin"
}

output "admin_role" {
  value = data.imagefactory_role.admin_role
}
