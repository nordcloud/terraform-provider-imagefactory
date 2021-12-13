// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

data "imagefactory_system_component" "hardening_level_1" {
  name = "Hardening level 1"
}

output "system_component" {
  value = data.imagefactory_system_component.hardening_level_1
}
