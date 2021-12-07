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

resource "imagefactory_account" "aws_account" {
  alias             = "IF AWS Account"
  description       = "Account to distribute AWS images"
  cloud_provider    = "AWS"
  cloud_provider_id = "822506164844"
  credentials {
    aws {
      role_arn         = "arn:aws:iam::822506164844:role/ImageFactoryMasterRole"
      role_external_id = "lieGhohY6ahv2aijieZ9"
    }
  }
}

output "aws_account" {
  value = imagefactory_account.aws_account
}
