package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	server "github.com/nii236/nii-finance/tickSubscriber/server"
)

var (
	natsURL  = "192.168.99.100:32773"
	log      *logrus.Logger
	settings = settingsContainer{}
)

type settingsContainer struct {
	Pair string
}

func getOptions(o *micro.Options) {
	o.Server = server.NewNatsServer(settings.Pair)
}

func main() {
	log = logrus.New()
	service := micro.NewService(
		micro.Flags(cli.StringFlag{
			Name:   "pair",
			Value:  "AUDUSD",
			Usage:  "Select pair to subsribe",
			EnvVar: "PAIR",
		}),
		micro.Action(func(c *cli.Context) {
			settings.Pair = c.String("pair")
		}),
	)

	service.Init(
		micro.Name("PairSubscriber"),
		getOptions,
	)
	log.Infoln("Starting pair subscription service for:", settings.Pair)
	service.Run()
}
