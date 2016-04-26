package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	server "github.com/nii236/nii-finance/tickSubscriber/server"
)

var (
	natsURL = "192.168.99.100:32773"
	log     *logrus.Logger
)

type settingsContainer struct {
	Pair string
}

func getOptions(o *micro.Options) {
	o.Server = server.NewNatsServer()
}

func main() {
	log = logrus.New()

	settings := settingsContainer{}

	service := micro.NewService(
		micro.Flags(cli.StringFlag{
			Name:   "pair",
			Value:  "AUDUSD",
			Usage:  "Select which pair to what you wish to subscribe",
			EnvVar: "PAIR",
		}),
		getOptions,
		micro.Action(func(c *cli.Context) {
			settings.Pair = c.String("pair")
		}),
	)
	service.Run()
}
