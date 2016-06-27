package client

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport"
	natsBroker "github.com/micro/go-plugins/broker/nats"
	natsRegistry "github.com/micro/go-plugins/registry/nats"
	natsTransport "github.com/micro/go-plugins/transport/nats"
)

var natsURL = "192.168.99.100:32773"
var log *logrus.Logger

func getOptions(o *client.Options) {
	o.Broker = newNatsBroker()
	o.Registry = newNatsRegistry()
	o.Transport = newNatsTransport()
}

func newNatsBroker() broker.Broker {
	return natsBroker.NewBroker(func(o *broker.Options) {
		o.Addrs = []string{natsURL}
	})
}

func newNatsRegistry() registry.Registry {
	return natsRegistry.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{natsURL}
	})
}

func newNatsTransport() transport.Transport {
	return natsTransport.NewTransport(func(o *transport.Options) {
		o.Addrs = []string{natsURL}
	})
}

func NewNatsClient(pair string, currency string) client.Client {
	log = logrus.New()
	log.Infoln("Starting up tickSubscriber...")
	log.Infoln("Pair:", pair)
	log.Infoln("Currency:", currency)
	cmd.Init()
	log = logrus.New()
	client.NewClient(getOptions)
	log.Info("Running tick subscriber client...")
	return client.NewClient(getOptions)
}
