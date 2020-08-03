# Server Block Resource

This resource represents a [Caddyfile](https://caddyserver.com/docs/caddyfile) in Caddy configuration directories.

## Example Usage

```hcl
# This will create file //etc/caddy/conf.d/test.conf
resource "caddy_server_block" "my-server" {
  filename = "test.conf"
  markers = {
    docker_port = docker_container.web.ports.external
    docker_ports = "${docker_container.web.ports.external},${docker_container.web2.ports.external}"
  }
  markers_split = {
    docker_ports = ","
  }
  content = <<EOF
# content of file here
# external docker port is: {# docker_port #}
# access web port in array: {# docker_ports[0] #}
# access web2 port in array: {# docker_ports[1] #}
EOF
}
```

## Argument Reference

* `filename` - (Required) Name of the configuration file
* `content` - (Required) Content of the configuration file
* `markers`- (Optional) Key-Value map. Keys specified as marker (e.g. `{# key #}`, `{~ key ~}`, `{* key *}`) will be replaced by the assigned value.