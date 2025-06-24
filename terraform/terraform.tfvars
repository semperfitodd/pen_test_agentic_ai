domain = "brewsentry.com"

ecr_repos = {
  front_end = "front_end"
  processor = "processor"
  scanner = "scanner"
}

eks_cluster_version = "1.33"

eks_node_instance_type = ["t3.medium"]

environment = "pen_test_ai"

region = "us-east-1"

vpc_cidr = "10.12.0.0/16"
