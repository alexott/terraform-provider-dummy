package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func provider() *schema.Provider {
	type providerConfig map[string]string
	withName := map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
	}
	return &schema.Provider{
		Schema: withName,
		ConfigureContextFunc: func(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
			return providerConfig{
				"name": rd.Get("name").(string),
			}, nil
		},
		ResourcesMap: map[string]*schema.Resource{
			"dummy_thing": {
				Schema: withName,
				CreateContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {
					conf := i.(providerConfig)
					rd.SetId(fmt.Sprintf("%s/%s", conf["name"], rd.Get("name").(string)))
					return nil
				},
				ReadContext: schema.NoopContext,
				UpdateContext: schema.NoopContext,
				DeleteContext: schema.NoopContext,
			},
		},
	}
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider,
	})
}
