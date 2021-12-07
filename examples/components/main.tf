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
  api_key = "8f4fa464-bbfb-4c47-879b-a841a57228d7/72a8da286c68822170b6179eed36cc29615f56d02e59c072"
  api_url = "https://api.imagefactory.dev.nordcloudapp.com/graphql"
}

data "imagefactory_system_components" "all" {}

output "all_system_components" {
  value = data.imagefactory_system_components.all.components
}
