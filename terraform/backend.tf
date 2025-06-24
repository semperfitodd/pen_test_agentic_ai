terraform {
  backend "s3" {
    bucket = "bsc.sandbox.terraform.state"
    key    = "pen_test_ai/terraform.tfstate"
    region = "us-east-2"

    use_lockfile = true
  }
}
