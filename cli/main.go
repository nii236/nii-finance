package main

import (
    "github.com/micro/go-micro/cmd"
    _ "github.com/micro/go-plugins/registry/nats"
)

func main() {
    cmd.Init()
}
