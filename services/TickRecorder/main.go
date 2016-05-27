package main

import (
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"log"
	"net"
	"open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/subscriber"
)

func opts(o *micro.Options) {
	o.Server = server.NewServer()
}

func handle() {
	log.Println("Tick received")
}

func main() {
	cmd.Init()
	log.Println("Starting up Tick Recorder...")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
	}
	log.Println("Interfaces:")
	for _, add := range addrs {
		log.Println(add.Network()+":", add.String())
	}
	s := micro.NewService(opts)
	if err := s.Server().Subscribe(
		server.NewSubscriber(
			"go.micro.srv.TickRecorder",
			new(subscriber.Tick),
			func(so *server.SubscriberOptions) {
				so.Queue = "tickqueue1"
			},
		),
	); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Println(err)
	}
}
