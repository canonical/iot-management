# Iot Management EKS

Deploys the IoT Management Service to EKS using Terraform.

## Prerequisites

To deploy this project, you'll need the following. 

- Used the Device Management EKS project to define the base configuration. See the [README.md](https://github.com/everactive/device-management-eks) for more information. 

## Install

This repo comes with a Dockerfile with all the tooling needed to deploy this to your AWS infrastructure. For building the tools container, see the [README.md](https://github.com/everactive/device-management-eks) for more information.

`make console/tools`

This will drop you into the container with Terraform and other tools needed to run the deployment. This also mounts your local `~/.aws` and `~/.kube` files to share credentials. 

## Configuration

You'll need to fill in the required input variables defined in the [Terraform Inputs](#inputs). The make commands expect that these variables are described in the [stacks folder](/stacks). For an example, see the [test-example.tfvars.json](stacks/test-example.tfvars.json). These variables will need to reference the shared Terraform `state_bucket`, `state_kms_alias`, and `state_dm_key` that was defined by the Device Management EKS project. See the [Terraform Docs](#Terraform Docs) below for more details.

## Usage

Assuming you completed the required steps in the [Device Management DB README.md](https://github.com/everactive/device-management-eks) you'll be able to run.

`make STACK_NAME=test-example deploy` 

This will deploy the management services to the EKS cluster you defined.


### Applying updates

To apply any updates pull this repo. Then run `make update` to pull the latest version of the submodules. Then run `make STACK_NAME=test-example deploy`. 

### Tearing it down

Run the command `make destroy`. This will prompt you for approval before applying the changes. 

## Terraform Docs

<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |
| <a name="provider_kubectl"></a> [kubectl](#provider\_kubectl) | n/a |
| <a name="provider_kubernetes"></a> [kubernetes](#provider\_kubernetes) | n/a |
| <a name="provider_null"></a> [null](#provider\_null) | n/a |
| <a name="provider_terraform"></a> [terraform](#provider\_terraform) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_route53_record.iot_management](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record) | resource |
| [kubectl_manifest.iot_management](https://registry.terraform.io/providers/gavinbunney/kubectl/latest/docs/resources/manifest) | resource |
| [null_resource.rollout](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) | resource |
| [aws_eks_cluster.cluster](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/eks_cluster) | data source |
| [aws_elb.iot_management](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/elb) | data source |
| [aws_kms_key.state](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/kms_key) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |
| [aws_route53_zone.domain](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone) | data source |
| [kubectl_file_documents.management_manifests](https://registry.terraform.io/providers/gavinbunney/kubectl/latest/docs/data-sources/file_documents) | data source |
| [kubernetes_service.iot_management](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/data-sources/service) | data source |
| [terraform_remote_state.dm](https://registry.terraform.io/providers/hashicorp/terraform/latest/docs/data-sources/remote_state) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_region"></a> [region](#input\_region) | AWS region | `string` | n/a | yes |
| <a name="input_state_bucket"></a> [state\_bucket](#input\_state\_bucket) | Terraform shared S3 Bucket | `string` | n/a | yes |
| <a name="input_state_dm_key"></a> [state\_dm\_key](#input\_state\_dm\_key) | Key to shared device management state | `string` | n/a | yes |
| <a name="input_state_key"></a> [state\_key](#input\_state\_key) | Terraform state key | `string` | n/a | yes |
| <a name="input_state_kms_alias"></a> [state\_kms\_alias](#input\_state\_kms\_alias) | KMS key used to encrypt Terraform state data | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_management_url"></a> [management\_url](#output\_management\_url) | n/a |
<!-- END_TF_DOCS -->