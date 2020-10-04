# Data source: sendgrid_whitelist_ip

Use this data source to retrieve information about a whitelist ip.

## Example Usage

```hcl
data "sendgrid_whitelist_ip" "first" {
  id = 1945952
}
```

## Argument Reference

The following arguments are supported:

- `id` - (Required) Rule ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `ip` - Whitelisted IP.
- `id` - Rule ID.
- `created_at` - The time when the rule was created.
- `updated_at` - The time when the rule was updated.
