# terraform-aws-vpc-with-subnets
A terraform module to create a vpc with dynamic subnets and VPC endpoints to allow for launching fargate tasks from private subnets, logging, and using SSM Parameter Store and KMS.

## Inputs:
| Variable Name          | Type            | Description                                       | Default                                    |
|------------------------|-----------------|---------------------------------------------------|--------------------------------------------|
| var.vpc_cidr_block     | String          | CIDR block to use for the VPC.                    | "172.16.0.0/16"                            |
| var.region             | String          | AWS Region to deploy in.                          | "us-east-1"                                |
| var.availability_zones | List of Strings | Availability Zones for the VPC.                   | ["us-east-1a", "us-east-1b", "us-east-1c"] |
| var.vpc_name           | String          | Name for the VPC.                                 |                                            |
## Outputs:
| Output Name                            | Description                                          |
|----------------------------------------|------------------------------------------------------|
| availability_zones                     | List of Availability Zones where subnets were created |
| availability_zone_ids                  | List of Availability Zones IDs where subnets were created, when available |
| public_subnet_ids                      | IDs of the created public subnets                   |
| public_subnet_arns                     | ARNs of the created public subnets                   |
| private_subnet_ids                     | IDs of the created private subnets                   |
| private_subnet_arns                    | ARNs of the created private subnets                   |
| public_subnet_cidrs                    | IPv4 CIDR blocks of the created public subnets       |
| private_subnet_cidrs                   | IPv4 CIDR blocks of the created private subnets      |
| public_route_table_ids                 | IDs of the created public route tables               |
| private_route_table_ids                | IDs of the created private route tables              |
| public_network_acl_id                  | ID of the Network ACL created for public subnets     |
| private_network_acl_id                 | ID of the Network ACL created for private subnets    |
| nat_gateway_ids                        | IDs of the NAT Gateways created                       |
| nat_instance_ids                       | IDs of the NAT Instances created                      |
| nat_instance_ami_id                    | ID of the AMI used by NAT instance                    |
| nat_ips                                | Elastic IP Addresses in use by NAT                    |
| nat_eip_allocation_ids                 | Elastic IP allocations in use by NAT                  |
| az_private_subnets_map                 | Map of AZ names to list of private subnet IDs in the AZs |
| az_public_subnets_map                  | Map of AZ names to list of public subnet IDs in the AZs |
| az_private_route_table_ids_map         | Map of AZ names to list of private route table IDs in the AZs |
| az_public_route_table_ids_map          | Map of AZ names to list of public route table IDs in the AZs |
| named_private_subnets_map              | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of private subnet IDs |
| named_public_subnets_map               | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of public subnet IDs |
| named_private_route_table_ids_map      | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of private route table IDs |
| named_public_route_table_ids_map       | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of public route table IDs |
| named_private_subnets_stats_map        | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of objects with each object having three items: AZ, private subnet ID, private route table ID |
| named_public_subnets_stats_map         | Map of subnet names (specified in `subnets_per_az_names` variable) to lists of objects with each object having three items: AZ, public subnet ID, public route table ID |
| vpc_id                                 | The ID of the VPC                                    |
| vpc_arn                                | The ARN of the VPC                                   |
| vpc_cidr_block                         | The primary IPv4 CIDR block of the VPC               |
| vpc_main_route_table_id                | The ID of the main route table associated with this VPC |
| vpc_default_network_acl_id             | The ID of the network ACL created by default on VPC creation |
| vpc_default_security_group_id          | The ID of the security group created by default on VPC creation |
| vpc_default_route_table_id             | The ID of the route table created by default on VPC creation |
| additional_cidr_blocks                 | A list of the additional IPv4 CIDR blocks associated with the VPC |
| additional_cidr_blocks_to_association_ids | A map of the additional IPv4 CIDR blocks to VPC CIDR association IDs |
| igw_id                                 | The ID of the Internet Gateway                        |
