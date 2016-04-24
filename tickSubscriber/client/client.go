package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport"
	natsBroker "github.com/micro/go-plugins/broker/nats"
	natsRegistry "github.com/micro/go-plugins/registry/nats"
	natsTransport "github.com/micro/go-plugins/transport/nats"
	"github.com/nii236/nii-forex/tickRecorder/proto"
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

// publishes a message
func pub(c client.Client) {
	msg := c.NewPublication("topic.go.micro.srv.example", &tick.Message{
		Say: "This is a publication",
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), metadata.MD{
		"X-User-Id": []string{"john"},
		"X-From-Id": []string{"script"},
	})

	// publish message
	if err := c.Publish(ctx, msg); err != nil {
		log.Infoln("pub err: ", err)
		return
	}

	log.Infoln("Published:", msg)
}

func NewNatsClient() client.Client {
	return client.NewClient(getOptions)
}

func main() {
	c := NewNatsClient()
	log = logrus.New()

	log.Info("Publishing message...")
	pub(c)
}
