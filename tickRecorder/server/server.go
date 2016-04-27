package server

import (
	"github.com/Sirupsen/logrus"
	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
	natsBroker "github.com/micro/go-plugins/broker/nats"
	natsRegistry "github.com/micro/go-plugins/registry/nats"
	natsTransport "github.com/micro/go-plugins/transport/nats"
	"github.com/nii236/nii-finance/tickRecorder/server/handler"
)

var log *logrus.Logger
var natsURL = "192.168.99.100:32773"

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
	influx, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr: "http://192.168.99.100:8086",
	})
	if err != nil {
		log.Errorln("Error connecting to Influx:", err)
	}

	s.Subscribe(s.NewSubscriber("go.micro.srv.tickStream", handler.NewTickHandler(log, influx)))

	return s
}
