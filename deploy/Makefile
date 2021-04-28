.PHONY: console/tools lint init deploy plan apply destroy destroy! state push outputs docs/tf

COMPONENT_NAME=iot-management
STACK_NAME=""
KMS_KEY_ALIAS=$(shell cat ./stacks/${STACK_NAME}.tfvars.json | jq -r .state_kms_alias)
BACKEND_KEY=$(shell cat ./stacks/${STACK_NAME}.tfvars.json | jq -r .state_key)
BUCKET_NAME=$(shell cat ./stacks/${STACK_NAME}.tfvars.json | jq -r .state_bucket)
REGION=$(shell cat ./stacks/${STACK_NAME}.tfvars.json | jq -r .region)
KMS_KEY=$(shell aws kms describe-key --key-id ${KMS_KEY_ALIAS} --query 'KeyMetadata.Arn' --output text)

lint:
	terraform fmt
	terraform validate
	tflint
init:
	terraform init \
	-reconfigure \
	-input=false \
	-backend-config "region=${REGION}" \
	-backend-config "bucket=${BUCKET_NAME}" \
	-backend-config "key=${BACKEND_KEY}" \
	-backend-config "kms_key_id=${KMS_KEY}"
plan:
	terraform plan -input=false \
	-var-file "stacks/${STACK_NAME}.tfvars.json"
apply:
	terraform apply -input=false -auto-approve \
	-var-file "stacks/${STACK_NAME}.tfvars.json"
state:
	terraform state list
destroy: 
	terraform destroy \
	-var-file "stacks/${STACK_NAME}.tfvars.json" 
destroy!:
	terraform destroy -auto-approve \
	-var-file "stacks/${STACK_NAME}.tfvars.json"
deploy: init apply
docs/tf:
	terraform-docs markdown table --output-file README.md --output-mode inject .
console/tools:
	docker run  \
	-v $(PWD):/home/device-manager \
	-v ~/.aws:/root/.aws \
	-v ~/.kube:/root/.kube \
	-it --rm everactive/devops-tools:latest /bin/bash