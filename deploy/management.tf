data "aws_route53_zone" "domain" {
  zone_id = local.hosted_zone
}

data "kubectl_file_documents" "management_manifests" {
  content = templatefile(
    "${path.module}/pods/k8s-management.yaml",
    { IMAGE        = "${local.docker_namespace}/iot-management:${local.docker_tag}"
      HOST         = "${local.manager_domain}.${data.aws_route53_zone.domain.name}"
      IP_WHITELIST = local.ip_whitelist_private
    CERT_ARN = local.cert_arn }
  )
}

resource "kubectl_manifest" "iot_management" {
  override_namespace = local.namespace
  count              = length(data.kubectl_file_documents.management_manifests.documents)
  yaml_body          = element(data.kubectl_file_documents.management_manifests.documents, count.index)
  wait               = true
}

data "kubernetes_service" "iot_management" {
  depends_on = [
    kubectl_manifest.iot_management
  ]
  metadata {
    name      = "management"
    namespace = local.namespace
  }
}

data "aws_elb" "iot_management" {
  depends_on = [
    kubectl_manifest.iot_management
  ]
  name = split("-", data.kubernetes_service.iot_management.status.0.load_balancer.0.ingress.0.hostname)[0]
}

resource "aws_route53_record" "iot_management" {
  zone_id = local.hosted_zone
  name    = local.manager_domain
  type    = "A"
  alias {
    name                   = data.aws_elb.iot_management.dns_name
    zone_id                = data.aws_elb.iot_management.zone_id
    evaluate_target_health = true
  }
}

resource "null_resource" "rollout" {
  depends_on = [
    aws_route53_record.iot_management
  ]
  triggers = {
    "timestamp" = timestamp()
  }
  provisioner "local-exec" {
    command = <<EOF
      aws eks update-kubeconfig --region ${local.region} --name ${local.eks_cluster};
      kubectl rollout restart deploy management -n=${local.namespace};
    EOF
  }
}