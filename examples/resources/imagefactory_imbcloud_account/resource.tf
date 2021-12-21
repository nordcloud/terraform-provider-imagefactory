// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_ibmcloud_account" "ibmcloud_account" {
  alias       = "IF IBMCLOUD Account"
  description = "IF IBMCLOUD Account to distribute IBMCLOUD images"
  account_id  = "1234567"
  access {
    apikey              = "APIKEY"
    region              = "eu-de"
    cos_bucket          = "nordcloudimagefactory-bucket"
    resource_group_name = "nordcloudimagefactory-rg"
    resource_group_id   = "asdfg12345zxcvb67890tyuiohjkllba"
  }
}

output "ibmcloud_account" {
  value = imagefactory_ibmcloud_account.ibmcloud_account
}
