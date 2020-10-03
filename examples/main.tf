variable "api_key" {}

terraform {
  required_providers {
    sendgrid = {
      versions = ["0.0.1"]
      source = "hashicorp.com/tatsuo48/sendgrid"
    }
  }
}

provider "sendgrid" {
  api_key = var.api_key
}

data "sendgrid_whitelist_ip" "first" {
  id = 1945952
}

resource "sendgrid_whitelist_ip" "first" {
  ip = "192.168.0.1/32"
}
