provider "azurerm" {
  version = "=1.29.0"
}

resource "azurerm_resource_group" "terraform-backend" {
  name      = "${var.resource_group}"
  location  = "${var.location}"
}

resource "azurerm_storage_account" "terraform-backend" {
  name                     = "${var.storage-account}"
  resource_group_name      = "${azurerm_resource_group.terraform-backend.name}"
  location                 = "${var.location}"
  account_tier             = "Standard"
  account_replication_type = "LRS"
}