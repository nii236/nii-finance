package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/examples/server/subscriber"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
)

var log *logrus.Logger

func getOptions(o *server.Options) {
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

func NewNatsServer() server.Server {

	log = logrus.New()
	s := server.NewServer(getOptions)
	log.Infoln("New server", s.String())

	s.Subscribe(s.NewSubscriber("topic.go.micro.srv.example", subscriber.Handler))

	return s
}
