// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

data "imagefactory_api_key" "api_key" {
  name = "test_api_key"
}

resource "imagefactory_role_binding" "key_binding" {
  kind = "API_KEY"
  role = "READ_ONLY"
  subject = data.imagefactory_api_key.api_key.id
}
