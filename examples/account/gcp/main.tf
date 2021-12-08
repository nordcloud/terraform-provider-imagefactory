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

resource "imagefactory_gcp_project" "gcp_project" {
  alias       = "IF GCP Project"
  description = "GCP project to distribute GCP images"
  project_id  = "project-name-123456"
  access {
    type                        = "service_account"
    private_key                 = <<-EOT
      -----BEGIN PRIVATE KEY-----
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIVATEKEYPRIV
      PRIVATEKEYPRIVATEKEYPRIVATEKEY
      -----END PRIVATE KEY-----
    EOT
    auth_uri                    = "https://accounts.google.com/o/oauth2/auth"
    client_id                   = "123456789012345678901"
    client_email                = "project@project-name-123456.iam.gserviceaccount.com"
    token_uri                   = "https://oauth2.googleapis.com/token"
    auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
    client_x509_cert_url        = "https://www.googleapis.com/robot/v1/metadata/x509/project%40project-name-123456.iam.gserviceaccount.com"
    project_id                  = "project-name-123456"
    private_key_id              = "asdfghjklqwertyuio1234567890zxcvbnm12345"
  }
}

output "gcp_project" {
  value = imagefactory_gcp_project.gcp_project
}
