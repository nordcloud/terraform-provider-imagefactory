// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

data "imagefactory_custom_component" "custom_component" {
  name = "Install nginx"
  cloud_provider = "AWS"
  stage = "BUILD"
}

output "custom_component" {
  value = data.imagefactory_custom_component.custom_component
}
