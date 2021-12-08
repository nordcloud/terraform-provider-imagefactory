// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

terraform {
  required_version = ">= 0.14"
  required_providers {
    imagefactory = {
      source  = "nordcloud.com/mc/imagefactory"
      version = "~> 1.0"
    }
  }
}

provider "imagefactory" {
  api_key = "KEY"
  api_url = "https://api.imagefactory.dev.nordcloudapp.com/graphql"
}

data "imagefactory_system_component" "hardening-level-1" {
  name = "Hardening level 1"
}

output "component" {
  value = data.imagefactory_system_component.hardening-level-1
}
