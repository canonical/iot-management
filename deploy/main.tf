terraform {
  backend "s3" {
  }
  required_providers {
    kubectl = {
      source = "gavinbunney/kubectl"
    }
  }
}

data "aws_region" "current" {}

locals {
  namespace            = data.terraform_remote_state.dm.outputs.namespace
  eks_cluster          = data.terraform_remote_state.dm.outputs.eks_cluster
  hosted_zone          = data.terraform_remote_state.dm.outputs.hosted_zone
  mqtt_domain          = data.terraform_remote_state.dm.outputs.mqtt_domain
  docker_namespace     = data.terraform_remote_state.dm.outputs.docker_namespace
  docker_tag           = data.terraform_remote_state.dm.outputs.docker_tag
  region               = data.aws_region.current.name
  kms_key_arn          = data.aws_kms_key.state.arn
  state_s3_bucket      = var.state_bucket
  state_s3_key         = var.state_key
  state_kms_key_alias  = var.state_kms_alias
  state_dm_key         = var.state_dm_key
  ip_whitelist_public  = data.terraform_remote_state.dm.outputs.ip_whitelist_public
  ip_whitelist_private = data.terraform_remote_state.dm.outputs.ip_whitelist_private
  cert_arn             = data.terraform_remote_state.dm.outputs.cert_arn
  identity_domain      = data.terraform_remote_state.dm.outputs.identity_domain
  manager_domain       = data.terraform_remote_state.dm.outputs.manager_domain
}

data "aws_kms_key" "state" {
  key_id = local.state_kms_key_alias
}


data "terraform_remote_state" "dm" {
  backend = "s3"
  config = {
    bucket     = local.state_s3_bucket
    key        = local.state_dm_key
    region     = local.region
    kms_key_id = local.kms_key_arn
  }
}

data "aws_eks_cluster" "cluster" {
  name = local.eks_cluster
}

provider "kubectl" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
  exec {
    api_version = "client.authentication.k8s.io/v1alpha1"
    args        = ["eks", "get-token", "--cluster-name", local.eks_cluster]
    command     = "aws"
  }
  load_config_file = false
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
  exec {
    api_version = "client.authentication.k8s.io/v1alpha1"
    args        = ["eks", "get-token", "--cluster-name", local.eks_cluster]
    command     = "aws"
  }
}