output "management_url" {
  value = "https://${aws_route53_record.iot_management.fqdn}"
}
