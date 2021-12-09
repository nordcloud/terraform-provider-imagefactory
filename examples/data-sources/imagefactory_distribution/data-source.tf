// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

data "imagefactory_distribution" "ubuntu18" {
  name           = "Ubuntu Server 18.04 LTS"
  cloud_provider = "AWS"
}

output "distro" {
  value = data.imagefactory_distribution.ubuntu18
}
