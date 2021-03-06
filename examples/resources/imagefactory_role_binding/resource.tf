// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_role_binding" "user_binding" {
  kind = "USER"
  role = "ADMIN"
  subject = "user@nordcloud.com"
}

data "imagefactory_api_key" "api_key" {
  name = "test_api_key"
}

resource "imagefactory_role_binding" "key_binding" {
  kind = "API_KEY"
  role = "READ_ONLY"
  subject = data.imagefactory_api_key.api_key.id
}

output "user_binding" {
  value = imagefactory_role_binding.user_binding
}
