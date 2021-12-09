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

data "imagefactory_custom_component" "custom-component" {
  name = "custom component name"
  cloud_provider = "AWS"
  stage = "TEST"
}

output "custom-component" {
  value = data.imagefactory_custom_component.custom-component
}
