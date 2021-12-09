// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

data "imagefactory_system_component" "hardening-level-1" {
  name = "Hardening level 1"
  cloud_provider = "AWS"
  stage = "BUILD"
}

output "system-component" {
  value = data.imagefactory_system_component.hardening-level-1
}
