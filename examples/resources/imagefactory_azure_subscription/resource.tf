// Copyright 2021 Nordcloud Oy or its affiliates. All Rights Reserved.

resource "imagefactory_azure_subscription" "azure_subscription" {
  alias           = "IF Azure Subscription"
  description     = "Azure subscription to distribute Azure images"
  subscription_id = "12345678-9012-3456-7890-123456789012"
  access {
    resource_group_name  = "RG-NAME"
    tenant_id            = "12345678-9012-3456-7890-123456789012"
    app_id               = "12345678-9012-3456-7890-123456789012"
    password             = "PASSWORD"
    storage_account      = "STORAGEACCOUNT"
    storage_account_key  = "ACCOUNT_KEY"
    shared_image_gallery = "IMAGE_GALLERY"
  }
}

output "azure_subscription" {
  value = imagefactory_azure_subscription.azure_subscription
}
