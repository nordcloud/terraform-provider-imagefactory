// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_aws_china_account" "aws_china_account" {
  alias       = "IF AWS China Account"
  description = "Account to distribute AWS images"
  account_id  = "123456789012"
  access {
    AWS_ACCESS_KEY_ID     = "A...B"
    AWS_SECRET_ACCESS_KEY = "A..SECRET..B"
  }
}

output "aws_china_account" {
  value = imagefactory_aws_china_account.aws_account
}
