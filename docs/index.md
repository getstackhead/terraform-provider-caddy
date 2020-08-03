# Caddy Provider

This provider can be used to manage Caddy configurations.

## Installation

1. Download a binary from the [release section](https://github.com/getstackhead/terraform-caddy/releases).
2. Install it into the user plugins directory.
   * Windows: `%APPDATA%\terraform.d\plugins`
   * other OS: `~/.terraform.d/plugins`

## Example Usage

```hcl
resource "caddy_server_block" "my-server" {
  filename = "test.conf"
  content = <<EOF
my-server.com {
  respond "Hello, world!"
}
sub.my-server.com {
  respond "Hello, world!"
}
EOF
}
```

The file will be stored inside the configured folder.

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g. `alias` and `version`), the following arguments are supported in the Nginx provider block:

* `config_folder` - (Optional) Folder where all configurations are stored. Default: `/etc/caddy/conf.d`

## Resources

### [server_block](./resources/server_block.md)