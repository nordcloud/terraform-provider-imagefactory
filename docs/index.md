---
layout: ""
page_title: "Provider: Imagefactory"
description: |-
  The Imagefactory provider provides resources to interact with the Nordcloud Klarity Imagefactory API.
---

# Imagefactory Provider

The Imagefactory provider provides resources to interact with the Nordcloud Klarity Imagefactory API.

Check https://docs.imagefactory.nordcloudapp.com for detailed product documentation.

## Example Usage

```terraform
terraform {
  required_version = ">= 0.14"
  required_providers {
    imagefactory = {
      source  = "nordcloud.com/klarity/imagefactory"
      version = "~> 1.0.0"
    }
  }
}

provider "imagefactory" {
  api_key = "API_KEY_ID/API_KEY_SECRET"
  api_url = "https://api.imagefactory.nordcloudapp.com/graphql"
}
```
