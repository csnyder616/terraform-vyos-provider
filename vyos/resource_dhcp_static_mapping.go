package vyos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDhcpStaticMapping() *schema.Resource {
	return &schema.Resource {
		Schema: map[string]*schema.Schema{
			"shared_network_name": &schema.Schema {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet": &schema.Schema {
				Type: schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema {
				Type: schema.TypeString,
				Required: true,
			},
			"mac_address": &schema.Schema {
				Type: schema.TypeString,
				Required: true,
			},
			"ip_address": &schema.Schema {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}

}
