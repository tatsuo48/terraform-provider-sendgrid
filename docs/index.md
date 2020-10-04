# Terraform Sendgrid provider

The Terraform Sendgrid provider is used to interact with many resources supported by Sendgrid.
The provider needs to be configured with the proper credentials before it can be used.

## Example Usage

```hcl
# Configure the provider
provider "sendgrid" {
    api_key = "SG.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

# Create a resource sendgrid_whitelist_ip
resource "sendgrid_whitelist_ip" "first" {
  ip = "192.168.0.1/32"
}

# Create a data source sendgrid_whitelist_ip
data "sendgrid_whitelist_ip" "first" {
  id = 1945952
}
```

## Authentication

The Sendgrid provider offers a flexible means of providing credentials for authentication.
The following methods are supported, and explained below in this order:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding `api_key` to the Sendgrid provider block, you can configure `host` and `subuser` too.

Usage:

```hcl
provider "sendgrid" {
    api_key = "SG.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

### Environment variables

You can provide your credentials via `SENDGRID_API_KEY`.

```hcl
provider "sendgrid" {}
```

Usage:

```shell
export SENDGRID_API_KEY="SG.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
terraform plan
```

## Testing

Credentials must be provided via the `SENDGRID_API_KEY` environment variable in order to run acceptance tests.

## Datasources/Resources reference

### sendgrid_whitelist_ip resource

- [sendgrid_whitelist_ip](resources/sendgrid_whitelist_ip.md)

### sendgrid_whitelist_ip data source

- [sendgrid_whitelist_ip](data-sources/sendgrid_whitelist_ip.md)
