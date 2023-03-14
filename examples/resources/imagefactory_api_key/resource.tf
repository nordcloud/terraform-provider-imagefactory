// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_api_key" "api_key" {
  name = "IF API Key"
}

output "api_key" {
  value = imagefactory_api_key.api_key
}
