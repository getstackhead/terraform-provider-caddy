package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"stackhead.io/terraform-caddy-provider/src/caddy"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"config_folder": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "/etc/caddy/conf.d",
				Description: "Folder where all configurations are stored",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"caddy_server_block": resourceServerBlock(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := caddy.Config{
		ConfigFolder: d.Get("config_folder").(string),
	}

	return config, nil
}
