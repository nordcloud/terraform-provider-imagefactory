// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_variable" "variable_test" {
  name      = "TEST"
  value     = "TEST_VALUE"
}

resource "imagefactory_custom_component" "component_with_variable" {
  name            = "Echo"
  description     = "Echo variable"
  stage           = "BUILD"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["LINUX"]
  content {
    script             = <<-EOT
      echo "$TEST"
    EOT
    provisioner = "SHELL"
  }
}
