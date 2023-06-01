// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_exoscale_organization" "exoscale_organization" {
  alias             = "IF Exoscale Organization"
  description       = "Exoscale Organization to distribute Exoscale templates"
  organization_name = "Exoscale Organization"
  access {
    api_key    = "EXOAPIKEY"
    api_secret = "APISECRET"
  }
}

output "exoscale_organization" {
  value = imagefactory_exoscale_organization.exoscale_organization
}
