// Copyright 2021-2022 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_aws_china_account" "aws_china_account" {
  alias       = "IF AWS China Account"
  description = "Account to distribute AWS images"
  account_id  = "123456789012"
  access {
    aws_access_key_id     = "A...B"
    aws_secret_access_key = "A..SECRET..B"
  }
  properties {
    s3_bucket_name = "s3_bucket_name"
    region = "cn-north-1"
  }
}

output "aws_china_account" {
  value = imagefactory_aws_china_account.aws_account
}
