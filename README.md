# Terraform Provider Sendgrid

## supported resources

Current supported only `ip_access_management`  
[ip_access_management](https://sendgrid.com/docs/API_Reference/Web_API_v3/ip_access_management.html)

### data sources

- sendgrid_whitelist_ip

```terraform
data "sendgrid_whitelist_ip" "first" {
  id = 1945952
}
```

### resources

- sendgrid_whitelist_ip

```terraform
resource "sendgrid_whitelist_ip" "first" {
  ip = "192.168.0.1/32"
}
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
cd examples
terraform init && terraform apply
```
