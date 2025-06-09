<!-- BEGIN_TF_DOCS -->
# terraform-aws-vpc-with-subnets

A terraform module to create a vpc with dynamic subnets
and VPC endpoints to allow for launching fargate tasks
from private subnets, logging, and using SSM Parameter
Store and KMS.

## Example Usage

```hcl
provider "aws" {
  region = var.region
}
module "vpc" {
  source             = "app.terraform.io/quansight/vpc-with-subnets/aws"
  version            = "v0.0.1"
  region             = var.region
  vpc_cidr_block     = var.vpc_cidr_block
  availability_zones = var.availability_zones
  vpc_name           = var.vpc_name
}
```

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.99.1 |

## Resources

| Name | Type |
|------|------|
| [aws_security_group.endpoint_security_group](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |
| [aws_security_group_rule.vpc_endpoint_access_egress](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule) | resource |
| [aws_security_group_rule.vpc_endpoint_access_ingress](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule) | resource |
| [aws_vpc_endpoint.dkr](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) | resource |
| [aws_vpc_endpoint.ecr](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) | resource |
| [aws_vpc_endpoint.kms](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) | resource |
| [aws_vpc_endpoint.logs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) | resource |
| [aws_vpc_endpoint.s3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) | resource |
| [aws_vpc_endpoint.ssm](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_availability_zones"></a> [availability\_zones](#input\_availability\_zones) | Availability Zones for the VPC | `list(string)` | <pre>[<br/>  "us-east1a",<br/>  "us-east1b",<br/>  "us-east1c"<br/>]</pre> | no |
| <a name="input_region"></a> [region](#input\_region) | AWS Region to deploy in | `string` | `"us-east-1"` | no |
| <a name="input_vpc_cidr_block"></a> [vpc\_cidr\_block](#input\_vpc\_cidr\_block) | Cidr block to use for the VPC. | `string` | `"172.16.0.0/16"` | no |
| <a name="input_vpc_name"></a> [vpc\_name](#input\_vpc\_name) | Name for the VPC | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_additional_cidr_blocks"></a> [additional\_cidr\_blocks](#output\_additional\_cidr\_blocks) | A list of the additional IPv4 CIDR blocks associated with the VPC |
| <a name="output_additional_cidr_blocks_to_association_ids"></a> [additional\_cidr\_blocks\_to\_association\_ids](#output\_additional\_cidr\_blocks\_to\_association\_ids) | A map of the additional IPv4 CIDR blocks to VPC CIDR association IDs |
| <a name="output_availability_zone_ids"></a> [availability\_zone\_ids](#output\_availability\_zone\_ids) | List of Availability Zones IDs where subnets were created, when available |
| <a name="output_availability_zones"></a> [availability\_zones](#output\_availability\_zones) | List of Availability Zones where subnets were created |
| <a name="output_az_private_route_table_ids_map"></a> [az\_private\_route\_table\_ids\_map](#output\_az\_private\_route\_table\_ids\_map) | Map of AZ names to list of private route table IDs in the AZs |
| <a name="output_az_private_subnets_map"></a> [az\_private\_subnets\_map](#output\_az\_private\_subnets\_map) | Map of AZ names to list of private subnet IDs in the AZs |
| <a name="output_az_public_route_table_ids_map"></a> [az\_public\_route\_table\_ids\_map](#output\_az\_public\_route\_table\_ids\_map) | Map of AZ names to list of public route table IDs in the AZs |
| <a name="output_az_public_subnets_map"></a> [az\_public\_subnets\_map](#output\_az\_public\_subnets\_map) | Map of AZ names to list of public subnet IDs in the AZs |
| <a name="output_igw_id"></a> [igw\_id](#output\_igw\_id) | The ID of the Internet Gateway |
| <a name="output_named_private_route_table_ids_map"></a> [named\_private\_route\_table\_ids\_map](#output\_named\_private\_route\_table\_ids\_map) | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of private route table IDs |
| <a name="output_named_private_subnets_map"></a> [named\_private\_subnets\_map](#output\_named\_private\_subnets\_map) | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of private subnet IDs |
| <a name="output_named_private_subnets_stats_map"></a> [named\_private\_subnets\_stats\_map](#output\_named\_private\_subnets\_stats\_map) | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of objects with each object having three items: AZ, private subnet ID, private route table ID |
| <a name="output_named_public_route_table_ids_map"></a> [named\_public\_route\_table\_ids\_map](#output\_named\_public\_route\_table\_ids\_map) | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of public route table IDs |
| <a name="output_named_public_subnets_map"></a> [named\_public\_subnets\_map](#output\_named\_public\_subnets\_map) | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of public subnet IDs |
| <a name="output_named_public_subnets_stats_map"></a> [named\_public\_subnets\_stats\_map](#output\_named\_public\_subnets\_stats\_map) | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of objects with each object having three items: AZ, public subnet ID, public route table ID |
| <a name="output_nat_eip_allocation_ids"></a> [nat\_eip\_allocation\_ids](#output\_nat\_eip\_allocation\_ids) | Elastic IP allocations in use by NAT |
| <a name="output_nat_gateway_ids"></a> [nat\_gateway\_ids](#output\_nat\_gateway\_ids) | IDs of the NAT Gateways created |
| <a name="output_nat_instance_ami_id"></a> [nat\_instance\_ami\_id](#output\_nat\_instance\_ami\_id) | ID of AMI used by NAT instance |
| <a name="output_nat_instance_ids"></a> [nat\_instance\_ids](#output\_nat\_instance\_ids) | IDs of the NAT Instances created |
| <a name="output_nat_ips"></a> [nat\_ips](#output\_nat\_ips) | Elastic IP Addresses in use by NAT |
| <a name="output_private_network_acl_id"></a> [private\_network\_acl\_id](#output\_private\_network\_acl\_id) | ID of the Network ACL created for private subnets |
| <a name="output_private_route_table_ids"></a> [private\_route\_table\_ids](#output\_private\_route\_table\_ids) | IDs of the created private route tables |
| <a name="output_private_subnet_arns"></a> [private\_subnet\_arns](#output\_private\_subnet\_arns) | ARNs of the created private subnets |
| <a name="output_private_subnet_cidrs"></a> [private\_subnet\_cidrs](#output\_private\_subnet\_cidrs) | IPv4 CIDR blocks of the created private subnets |
| <a name="output_private_subnet_ids"></a> [private\_subnet\_ids](#output\_private\_subnet\_ids) | IDs of the created private subnets |
| <a name="output_public_network_acl_id"></a> [public\_network\_acl\_id](#output\_public\_network\_acl\_id) | ID of the Network ACL created for public subnets |
| <a name="output_public_route_table_ids"></a> [public\_route\_table\_ids](#output\_public\_route\_table\_ids) | IDs of the created public route tables |
| <a name="output_public_subnet_arns"></a> [public\_subnet\_arns](#output\_public\_subnet\_arns) | ARNs of the created public subnets |
| <a name="output_public_subnet_cidrs"></a> [public\_subnet\_cidrs](#output\_public\_subnet\_cidrs) | IPv4 CIDR blocks of the created public subnets |
| <a name="output_public_subnet_ids"></a> [public\_subnet\_ids](#output\_public\_subnet\_ids) | IDs of the created public subnets |
| <a name="output_vpc_arn"></a> [vpc\_arn](#output\_vpc\_arn) | The ARN of the VPC |
| <a name="output_vpc_cidr_block"></a> [vpc\_cidr\_block](#output\_vpc\_cidr\_block) | The primary IPv4 CIDR block of the VPC |
| <a name="output_vpc_default_network_acl_id"></a> [vpc\_default\_network\_acl\_id](#output\_vpc\_default\_network\_acl\_id) | The ID of the network ACL created by default on VPC creation |
| <a name="output_vpc_default_route_table_id"></a> [vpc\_default\_route\_table\_id](#output\_vpc\_default\_route\_table\_id) | The ID of the route table created by default on VPC creation |
| <a name="output_vpc_default_security_group_id"></a> [vpc\_default\_security\_group\_id](#output\_vpc\_default\_security\_group\_id) | The ID of the security group created by default on VPC creation |
| <a name="output_vpc_endpoint_sg_id"></a> [vpc\_endpoint\_sg\_id](#output\_vpc\_endpoint\_sg\_id) | ID of the security group that allows access to the vpc endpoints |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | The ID of the VPC |
| <a name="output_vpc_main_route_table_id"></a> [vpc\_main\_route\_table\_id](#output\_vpc\_main\_route\_table\_id) | The ID of the main route table associated with this VPC |
<!-- END_TF_DOCS -->