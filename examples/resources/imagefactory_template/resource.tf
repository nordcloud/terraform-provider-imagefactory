# AWS Template

data "imagefactory_distribution" "ubuntu18" {
  name           = "Ubuntu Server 18.04 LTS"
  cloud_provider = "AWS"
}

resource "imagefactory_template" "template" {
  name            = "Ubuntu1804"
  description     = "Ubuntu 18.04 on AWS"
  cloud_provider  = "AWS"
  distribution_id = data.imagefactory_distribution.ubuntu18.id
  config {
    aws {
      region = "eu-west-1"
    }
    notifications {
      type = "SNS"
      uri  = "arn:aws:sns:eu-west-1:123456789012:Topic"
    }
    tags {
      key   = "KEY_ONE"
      value = "VALUE_A"
    }
    tags {
      key   = "KEY_TWO"
      value = "VALUE_B"
    }
  }
}

output "template" {
  value = imagefactory_template.template
}

# AZURE Template

data "imagefactory_distribution" "ubuntu18" {
  name           = "Ubuntu Server 18.04 LTS"
  cloud_provider = "AZURE"
}

resource "imagefactory_template" "template" {
  name            = "Ubuntu1804"
  description     = "Ubuntu 18.04 on Azure"
  cloud_provider  = "AZURE"
  distribution_id = data.imagefactory_distribution.ubuntu18.id
  config {
    azure {
      exclude_from_latest = true
      replica_regions     = ["westeurope"]
    }
    notifications {
      type = "SNS"
      uri  = "arn:aws:sns:eu-west-1:123456789012:Topic"
    }
    tags {
      key   = "KEY_ONE"
      value = "VALUE_A"
    }
    tags {
      key   = "KEY_TWO"
      value = "VALUE_B"
    }
  }
}

output "template" {
  value = imagefactory_template.template
}
