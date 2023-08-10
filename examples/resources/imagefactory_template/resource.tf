// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_custom_component" "build_template" {
  name            = "Install nginx"
  description     = "Install nginx on Ubuntu"
  stage           = "BUILD"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["LINUX"]
  content {
    script      = <<-EOT
      apt-get update && apt-get install nginx -y
    EOT
    provisioner = "SHELL"
  }
}

resource "imagefactory_custom_component" "test_component" {
  name            = "Test nginx"
  description     = "Test nginx is installed"
  stage           = "TEST"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["LINUX"]
  content {
    script      = <<-EOT
      ps aux | grep nginx
      systemctl is-active --quiet nginx || echo "nginx is not running"; exit 1
    EOT
    provisioner = "SHELL"
  }
}

data "imagefactory_system_component" "hardening-level-1" {
  name           = "Hardening level 1"
  cloud_provider = "AWS"
  stage          = "BUILD"
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
      id = imagefactory_custom_component.build_template.id
    }
    test_components {
      id = imagefactory_custom_component.test_component.id
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

# AZURE Template

data "imagefactory_distribution" "ubuntu18" {
  name           = "Ubuntu Server 18.04 LTS"
  cloud_provider = "AZURE"
}

resource "imagefactory_template" "azure_template" {
  name            = "Ubuntu1804"
  description     = "Ubuntu 18.04 on Azure"
  cloud_provider  = "AZURE"
  distribution_id = data.imagefactory_distribution.ubuntu18.id
  config {
    azure {
      exclude_from_latest = true
      eol_date_option     = true
      replica_regions     = ["westeurope"]
      vm_image_definition {
        name  = "Ubuntu1804"
        offer = "ubuntu-18_04-lts"
        sku   = "v1"
      }
    }
    notifications {
      type = "WEB_HOOK"
      uri  = "https://webhook.call.api.address"
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

output "azure_template" {
  value = imagefactory_template.azure_template
}

# AWS CHINA template

resource "imagefactory_template" "aws_china_template" {
  name            = "Ubuntu1804"
  description     = "Ubuntu 18.04 on AWS"
  cloud_provider  = "AWS"
  distribution_id = data.imagefactory_distribution.ubuntu18.id
  config {
    scope = "CHINA"
  }
}

# AWS template - copy image to selected account

resource "imagefactory_aws_account" "aws_account" {
  alias       = "IF AWS Account"
  description = "Account to distribute AWS images"
  account_id  = "123456789012"
  access {
    role_arn         = "arn:aws:iam::123456789012:role/ImageFactoryMasterRole"
    role_external_id = "lieGhohY6ahv2aijieZ9"
  }
}

data "imagefactory_distribution" "ubuntu18" {
  name           = "Ubuntu Server 18.04 LTS"
  cloud_provider = "AWS"
}

resource "imagefactory_template" "aws_template" {
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
    cloud_account_ids = [imagefactory_aws_account.aws_account.id]
  }
}

# AWS template - additional EBS volumes

data "imagefactory_distribution" "ubuntu18" {
  name           = "Ubuntu Server 18.04 LTS"
  cloud_provider = "AWS"
}

resource "imagefactory_template" "aws_template" {
  name            = "Ubuntu1804"
  description     = "Ubuntu 18.04 on AWS"
  cloud_provider  = "AWS"
  distribution_id = data.imagefactory_distribution.ubuntu18.id
  config {
    aws {
      region = "eu-west-1"
      additional_ebs_volumes {
        size        = 1
        device_name = "/dev/sdb"
      }
    }
    build_components {
      id = data.imagefactory_system_component.hardening-level-1.id
    }
  }
}

# EXOSCALE Template

data "imagefactory_distribution" "ubuntu22" {
  name           = "Ubuntu Server 22.04 LTS"
  cloud_provider = "EXOSCALE"
}

resource "imagefactory_template" "exoscale_template" {
  name            = "Ubuntu2204"
  description     = "Ubuntu Server 22.04 on Exoscale"
  cloud_provider  = "EXOSCALE"
  distribution_id = data.imagefactory_distribution.ubuntu22.id
  config {
    exoscale {
      zone = "de-fra-1"
    }
    notifications {
      type = "WEB_HOOK"
      uri  = "https://webhook.call.api.address"
    }
  }
}

output "template" {
  value = imagefactory_template.exoscale_template
}
