package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/nii236/nii-finance/tickRecorder/proto"
	"github.com/nii236/nii-finance/tickSubscriber/client"
	"golang.org/x/net/context"
)

var (
	natsURL  = "192.168.99.100:32773"
	log      *logrus.Logger
	settings = settingsContainer{}
	service  micro.Service
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

	service = micro.NewService(
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

func pub(time int64, bid float64, ask float64, last float64) {
	tick := &tickRecorder.Tick{
		Time: time,
		Bid:  bid,
		Ask:  ask,
		Last: last,
	}
	p := service.Client().NewPublication("go.micro.srv.tickStream", tick)
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "nii236",
		"X-From-Id": "tickSubscriber",
	})
	service.Client().Publish(ctx, p)

}
