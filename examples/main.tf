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

variable "template_name" {
  type    = string
  default = ""
}

data "imagefactory_distributions" "all" {}

output "all_distributions" {
  value = data.imagefactory_distributions.all.distributions
}

// output "template" {
//   value = {
//     for template in data.imagefactory_templates.all.templates :
//     template.id => template
//     if template.name == var.template_name
//   }
// }
