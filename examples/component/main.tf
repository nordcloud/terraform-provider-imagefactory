// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

terraform {
  required_version = ">= 0.14"
  required_providers {
    imagefactory = {
      source  = "nordcloud.com/klarity/imagefactory"
      version = "~> 1.0"
    }
  }
}

provider "imagefactory" {
  api_key = "KEY"
  api_url = "https://api.imagefactory.nordcloudapp.com/graphql"
}

resource "imagefactory_component" "component" {
  name            = "Install nginx"
  description     = "Install nginx on Ubuntu"
  stage           = "BUILD"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["LINUX"]
  content {
    script             = <<-EOT
      apt-get update && apt-get install nginx -y
    EOT
    provisioner = "SHELL"
  }
}

output "component" {
  value = imagefactory_component.component
}
