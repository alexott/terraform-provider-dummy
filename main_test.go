package main

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccResourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"dummy": func() (*schema.Provider, error) {
				return provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `
				provider "dummy" {
					name = "foo"
				}
				
				resource "dummy_thing" "this" {
					name = "bar"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"dummy_thing.this", "id", "foo/bar"),
				),
			},
		},
	})
}