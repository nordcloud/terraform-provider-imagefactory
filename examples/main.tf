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

data "imagefactory_distribution" "windows2019" {
  name = "Windows Server 2019"
  cloud_provider = "AWS"
}

output "distro" {
  value = data.imagefactory_distribution.windows2019
}

data "imagefactory_distributions" "all" {}

output "all_distributions" {
  value = data.imagefactory_distributions.all.distributions
}

data "imagefactory_distribution" "ubuntu18" {
  name           = "Ubuntu Server 18.04 LTS"
  cloud_provider = "AWS"
}

output "distro" {
  value = data.imagefactory_distribution.ubuntu18.id
}

resource "imagefactory_template" "template" {
  name            = "Ubuntu1804"
  description     = "dwadwalula"
  cloud_provider  = "AWS"
  distribution_id = data.imagefactory_distribution.ubuntu18.id
  config {
    aws {
      region = "eu-west-1"
    }
  }
}

output "template" {
  value = imagefactory_template.template
}
