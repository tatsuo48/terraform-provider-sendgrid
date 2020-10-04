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
  id = 1945952
}

resource "sendgrid_whitelist_ip" "first" {
  ip = "192.168.0.1/32"
}
