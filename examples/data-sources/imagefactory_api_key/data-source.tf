// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

data "imagefactory_api_key" "api_key" {
  name = "test_api_key"
}

resource "imagefactory_role_binding" "key_binding" {
  kind = "API_KEY"
  role_id = "12345678-9012-3456-7890-123456789012"
  subject = data.imagefactory_api_key.api_key.id
}
