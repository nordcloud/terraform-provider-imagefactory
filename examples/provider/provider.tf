/**
 * Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.
 */

terraform {
  required_version = ">= 0.14"
  required_providers {
    imagefactory = {
      source  = "nordcloud/imagefactory"
      version = "1.3.1"
    }
  }
}

provider "imagefactory" {
  api_key = "API_KEY_ID/API_KEY_SECRET"
  api_url = "https://api.imagefactory.nordcloudapp.com/graphql"
}
