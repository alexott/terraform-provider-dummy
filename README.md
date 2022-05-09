# terraform-provider-dummy

It cannot do much at all. Only a `dummy_thing` that does nothing at all.

```hcl
provider "dummy" {
    name = "foo"
}

resource "dummy_thing" "this" {
    name = "bar"
}

resource "another_dummy_thing" "that" {
    name = "baz"
}
```