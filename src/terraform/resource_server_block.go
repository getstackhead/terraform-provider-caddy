package terraform

import (
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"stackhead.io/terraform-caddy-provider/src/caddy"
)

func resourceServerBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerBlockCreate,
		Read:   resourceServerBlockRead,
		Update: resourceServerBlockUpdate,
		Delete: resourceServerBlockDelete,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the configuration file",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Content of the configuration file",
			},
			"markers": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Markers in content that should be replaced",
			},
			"markers_split": {
				Type:        schema.TypeMap,
				Default:     "",
				Description: "Define marker name as key and the character where the string is split as value. Chunks can be accessed as array",
				Optional:    true,
			},
		},
	}
}

func resourceServerBlockCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(caddy.Config)

	// Create file
	content := d.Get("content").(string)
	fullPathAvailable, err := caddy.CreateOrUpdateServerBlock(d.Get("filename").(string), content, config, d.Get("markers").(map[string]interface{}), d.Get("markers_split").(map[string]interface{}))
	if err != nil {
		return err
	}

	d.SetId(fullPathAvailable)
	return resourceServerBlockRead(d, m)
}

func resourceServerBlockRead(d *schema.ResourceData, m interface{}) error {
	content, err := caddy.ReadFile(d.Id())
	if err != nil {
		return err
	}
	d.Set("filename", filepath.Base(d.Id()))
	d.Set("content", content)
	return nil
}

func resourceServerBlockUpdate(d *schema.ResourceData, m interface{}) error {
	// Content changed: replace old file content
	if d.HasChange("content") || d.HasChange("variables") {
		_, err := caddy.CreateOrUpdateServerBlock(d.Id(), d.Get("content").(string), m.(caddy.Config), d.Get("markers").(map[string]interface{}), d.Get("markers_split").(map[string]interface{}))
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceServerBlockDelete(d *schema.ResourceData, m interface{}) error {
	if err := caddy.RemoveServerBlock(d.Get("filename").(string), m.(caddy.Config)); err != nil {
		return err
	}
	d.SetId("")
	return nil
}
