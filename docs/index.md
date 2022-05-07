# Dummy provider

It cannot do much at all. Only a `dummy_thing` that does nothing at all :)

```hcl
provider "dummy" {
    name = "foo"
}

resource "dummy_thing" "this" {
    name = "bar"
}
```

# Support

`dummy` provider is not supported at all and created for exploratory purposes only.