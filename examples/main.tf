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
  api_url = "https://api.hbi.dev.nordcloudapp.com/graphql"
}

variable "distribution_name" {
  type    = string
  default = "Windows Server 2019"
}

variable "distribution_provider" {
  type    = string
  default = "GCP"
}

data "imagefactory_distributions" "all" {}

// output "all_distributions" {
//   value = data.imagefactory_distributions.all.distributions
// }

output "distribution" {
  value = {
    for distribution in data.imagefactory_distributions.all.distributions :
    distribution.id => distribution
    if (
      distribution.name == var.distribution_name && distribution.provider == var.distribution_provider
    )
  }
}
