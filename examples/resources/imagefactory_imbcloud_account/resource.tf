// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_imbcloud_account" "imbcloud_account" {
  alias       = "IF IBMCloud Account"
  description = "IF IBMCloud Account to distribute IBMCloud images"
  account_id  = "1234567"
  access {
    apikey              = "APIKEY"
    region              = "eu-de"
    cos_bucket          = "nordcloudimagefactory-bucket"
    resource_group_name = "nordcloudimagefactory-rg"
    resource_group_id   = "asdfg12345zxcvb67890tyuiohjkllba"
  }
}

output "imbcloud_account" {
  value = imagefactory_imbcloud_account.imbcloud_account
}
