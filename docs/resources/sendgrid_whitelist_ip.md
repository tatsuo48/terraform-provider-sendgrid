# Resource: sendgrid_whitelist_ip

Creates an SendGrid whitelist IP.
For information about whitelist IP, see the [SendGrid ip-access-management Guide](https://sendgrid.com/docs/ui/account-and-settings/ip-access-management/).
API Reference is [here](https://sendgrid.com/docs/API_Reference/Web_API_v3/ip_access_management.html)

## Example Usage

```hcl
resource "sendgrid_whitelist_ip" "first" {
  ip = "192.168.0.1/32"
}
```

## Argument Reference

The following arguments are supported:

- `ip` - Whitelisted IP.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `ip` - Whitelisted IP.
- `id` - Rule ID.
- `created_at` - The time when the rule was created.
- `updated_at` - The time when the rule was updated.
