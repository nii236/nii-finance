package main

import (
	ccli "github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/micro/web"

	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
)

func main() {
	app := cmd.App()
	app.Commands = append(app.Commands, web.Commands()...)
	app.Action = func(context *ccli.Context) { ccli.ShowAppHelp(context) }
	cmd.Init(
		cmd.Name("micro"),
		cmd.Description("This version of micro contains only the web UI"),
		cmd.Version("latest"),
	)
}
