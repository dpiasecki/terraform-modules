variable "resource_group" {
  description = "The name of the resource group in which to create the container instance and Cosmos DB instance."
  default     = "terraform-state"
}

variable "location" {
  description = "The location for the resource group in which to create the container instance and Cosmos DB instance."
  default     = "eastus"
}

variable "storage-account" {
  description = "Name of the Azure storage account for Terraform state."
  default     = "nepetersstate"
}