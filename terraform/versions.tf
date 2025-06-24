provider "aws" {
  region = var.region

  default_tags {
    tags = {
      Project     = "pen_test_ai"
      Owners      = "Todd"
      Provisioner = "Terraform"
    }
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.96.0"
    }
  }
  required_version = "1.11.4"
}