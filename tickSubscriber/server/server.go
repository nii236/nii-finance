package server

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
	natsBroker "github.com/micro/go-plugins/broker/nats"
	natsRegistry "github.com/micro/go-plugins/registry/nats"
	natsTransport "github.com/micro/go-plugins/transport/nats"
)

var (
	log *logrus.Logger

	// sigChan receives os signals.
	sigChan = make(chan os.Signal, 1)
	// complete is used to report processing is done.
	complete = make(chan error)
	// shutdown provides system wide notification.
	shutdown = make(chan struct{})
	refresh  = make(chan bool)
	natsURL  = "192.168.99.100:32773"
)

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

// NewNatsServer will return a microservice instance to be used by go-micro's service.Run() func
func NewNatsServer(pair string) server.Server {
	cmd.App()
	log = logrus.New()
	s := server.NewServer(getOptions)
	return s
}
