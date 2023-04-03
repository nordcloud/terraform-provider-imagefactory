// Copyright 2023 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_api_key" "api_key" {
  name = "IF API Key"
}

output "api_key" {
  value = imagefactory_api_key.api_key
}

# Create API key with expiration date

resource "imagefactory_api_key" "api_key_with_expiration" {
  name       = "IF API Key"
  expires_at = "2023-04-27"
}

output "api_key_with_expiration" {
  value = imagefactory_api_key.api_key_with_expiration
}
