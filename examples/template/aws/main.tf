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

data "imagefactory_distribution" "ubuntu18" {
  name           = "Ubuntu Server 18.04 LTS"
  cloud_provider = "AWS"
}

resource "imagefactory_template" "template" {
  name            = "Ubuntu1804"
  description     = "Ubuntu 18.04 on AWS"
  cloud_provider  = "AWS"
  distribution_id = data.imagefactory_distribution.ubuntu18.id
  config {
    aws {
      region = "eu-west-1"
    }
    notifications {
      type = "SNS"
      uri  = "arn:aws:sns:eu-west-1:123456789012:Topic"
    }
    tags {
      key   = "KEY_ONE"
      value = "VALUE_A"
    }
    tags {
      key   = "KEY_TWO"
      value = "VALUE_B"
    }
  }
}

output "template" {
  value = imagefactory_template.template
}