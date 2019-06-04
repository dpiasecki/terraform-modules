output "blob_endpoint" {
  value = "${azurerm_storage_account.terraform-backend.primary_blob_endpoint}"
}