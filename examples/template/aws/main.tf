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

resource "imagefactory_component" "build_template" {
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

resource "imagefactory_component" "test_component" {
  name            = "Test nginx"
  description     = "Test nginx is installed"
  stage           = "TEST"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["LINUX"]
  content {
    script             = <<-EOT
      ps aux | grep nginx
      systemctl is-active --quiet nginx || echo "nginx is not running"; exit 1
    EOT
    provisioner = "SHELL"
  }
}

data "imagefactory_system_component" "hardening-level-1" {
  name = "Hardening level 1"
  cloud_provider = "AWS"
  stage = "BUILD"
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
    build_components {
      id = data.imagefactory_system_component.hardening-level-1.id
    }
    build_components {
      id = imagefactory_component.build_template.id
    }
    test_components {
      id = imagefactory_component.test_component.id
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
