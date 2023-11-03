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