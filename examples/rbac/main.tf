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

resource "imagefactory_role_binding" "user_binding" {
  kind = "USER"
  role = "ADMIN"
  subject = "kamil.piotrowski@nordcloud.com"
}

data "imagefactory_api_key" "api_key" {
  name = "test_api_key"
}

resource "imagefactory_role_binding" "key_binding" {
  kind = "API_KEY"
  role = "READ_ONLY"
  subject = data.imagefactory_api_key.api_key.id
}

output "user_binding" {
  value = imagefactory_role_binding.user_binding
}
