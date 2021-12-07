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

resource "imagefactory_aws_account" "aws_account" {
  alias       = "IF AWS Account"
  description = "Account to distribute AWS images"
  account_id  = "123456789012"
  access {
    role_arn         = "arn:aws:iam::123456789012:role/ImageFactoryMasterRole"
    role_external_id = "lieGhohY6ahv2aijieZ9"
  }
}

output "aws_account" {
  value = imagefactory_aws_account.aws_account
}
