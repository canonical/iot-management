provider "aws" {
  region = var.region
}

variable "region" {
  description = "AWS region"
  type        = string
}

variable "state_bucket" {
  description = "Terraform shared S3 Bucket"
  type        = string
}

variable "state_key" {
  description = "Terraform state key"
  type        = string
}

variable "state_kms_alias" {
  description = "KMS key used to encrypt Terraform state data"
  type        = string
}

variable "state_dm_key" {
  description = "Key to shared device management state"
  type        = string
}