package server

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
	natsBroker "github.com/micro/go-plugins/broker/nats"
	natsRegistry "github.com/micro/go-plugins/registry/nats"
	natsTransport "github.com/micro/go-plugins/transport/nats"
	"github.com/nii236/nii-forex/tickRecorder/server/subscriber"
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

func NewNatsServer() server.Server {

	log = logrus.New()
	s := server.NewServer(getOptions)
	log.Infoln("New server", s.String())

	s.Subscribe(s.NewSubscriber("topic.go.micro.srv.example", subscriber.Handler))

	return s
}
