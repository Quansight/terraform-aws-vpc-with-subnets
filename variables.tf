variable "vpc_cidr_block" {
    type = string
    description = "Cidr block to use for the VPC."
    default = "172.16.0.0/16"
}

variable "region" {
    type = string
    description = "AWS Region to deploy in"
    default = "us-east-1"
}

variable "availability_zones" {
    type = list(string)
    description = "Availability Zones for the VPC"
    default = [ "${var.region}a", "${var.region}b", "${var.region}c" ]
}

variable "vpc_name" {
    type = string
    description = "Name for the VPC"
}