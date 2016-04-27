package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/nii236/nii-finance/tickSubscriber/client"
)

var (
	natsURL  = "192.168.99.100:32773"
	log      *logrus.Logger
	settings = settingsContainer{}
)

type settingsContainer struct {
	pair     string
	currency string
}

func getOptions(o *micro.Options) {
	o.Client = client.NewNatsClient(settings.pair, settings.currency)
}

func main() {
	cmd.Init()
	log = logrus.New()

	service := micro.NewService(
		micro.Flags(cli.StringFlag{
			Name:   "pair",
			Value:  "USD",
			Usage:  "Select pair to subscribe",
			EnvVar: "PAIR",
		}),
		micro.Flags(cli.StringFlag{
			Name:   "currency",
			Value:  "JPY",
			Usage:  "Select pair to subscribe",
			EnvVar: "CURRENCY",
		}),
		micro.Action(func(c *cli.Context) {
			settings.pair = c.String("pair")
			settings.currency = c.String("currency")
		}),
		micro.Name("PairSubscriber"),
	)
	service.Init(getOptions)
	go exec(settings.pair, settings.currency)
	select {}
}
