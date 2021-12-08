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
  api_url = "https://api.imagefactory.nordcloudapp.com/graphql"
}

resource "imagefactory_imbcloud_account" "imbcloud_account" {
  alias       = "IF IBMCloud Account change"
  description = "IF IBMCloud Account to distribute IBMCloud images"
  account_id  = "1234567"
  access {
    apikey              = "APIKEY"
    region              = "eu-de"
    cos_bucket          = "nordcloudimagefactory-bucket"
    resource_group_name = "nordcloudimagefactory-rg"
    resource_group_id   = "asdfg12345zxcvb67890tyuiohjkllba"
  }
}

output "imbcloud_account" {
  value = imagefactory_imbcloud_account.imbcloud_account
}
