# mkdir -p .terraform/providers/registry.terraform.io/nfx01/dummy/0.0.1/darwin_amd64
# ln -s ../terraform-provider-dummy .terraform/providers/registry.terraform.io/nfx01/dummy/0.0.1/darwin_amd64/terraform-provider-dummy_v0.0.1
terraform {
  required_providers {
    dummy = {
      source = "nfx01/dummy"
      version = "0.0.1"
    }
  }
}

provider "dummy" {
    name = "foo"
}

resource "dummy_thing" "this" {
    name = "bar"
}
