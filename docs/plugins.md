# Plugin Development Guide

Plugins must expose:
- `Name() string`
- `Init()`

You can compile them with `go build -buildmode=plugin`.
Place `.so` files in `/plugins`.