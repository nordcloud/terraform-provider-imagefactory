// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_aws_account" "aws_account" {
  alias       = "IF AWS Account"
  description = "Account to distribute AWS images"
  account_id  = "123456789012"
  access {
    role_arn         = "arn:aws:iam::123456789012:role/ImageFactoryMasterRole"
    role_external_id = "lieGhohY6ahv2aijieZ9"
  }
}

output "aws_account" {
  value = imagefactory_aws_account.aws_account
}
