terraform {
  required_providers {
    sendgrid = {
      source = "hashicorp.com/tatsuo48/sendgrid"
    }
  }
}

provider "sendgrid" {
}

data "sendgrid_whitelist_ip" "first" {
  rule_id = 1234567
}

resource "sendgrid_whitelist_ip" "first" {
  ip = "127.0.0.1/32"
}
