---
layout: ""
page_title: "Provider: Imagefactory"
description: |-
  The Imagefactory provider provides resources to interact with the Nordcloud Klarity ImageFactory API.
---

# ImageFactory Provider

The Imagefactory provider provides resources to interact with the Nordcloud Klarity ImageFactory API.

## Nordcloud ImageFactory

Based on Nordcloud's extensive experience of managing cloud services, ImageFactory is the bomb-proof SaaS solution for fully managing image-hardening for multicloud.

ImageFactory offers:

- a hardening solution which is based on the most relevant and highly regarded security standards
- support for AWS, Azure, GCP, IBM Cloud and VMware, along with most common Windows and Linux versions
- automatic delivery updated images to all your cloud accounts and subscriptions
- various options to customize images with additional tools and extensions

Please check https://klarity.nordcloud.com to learn more about the ImageFactory and other Nordcloud products.

## Example Usage

```terraform
/**
 * Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.
 */

terraform {
  required_version = ">= 0.14"
  required_providers {
    imagefactory = {
      source  = "nordcloud/imagefactory"
      version = "1.5.0"
    }
  }
}

provider "imagefactory" {
  api_key = "API_KEY_ID/API_KEY_SECRET"
  api_url = "https://api.imagefactory.nordcloudapp.com/graphql"
}
```
