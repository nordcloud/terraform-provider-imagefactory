/**
 * Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.
 */

terraform {
  required_version = ">= 0.14"
  required_providers {
    imagefactory = {
      source  = "nordcloud/imagefactory"
      version = "1.4.2"
    }
  }
}

