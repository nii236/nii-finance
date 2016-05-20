package main

import (
	"log"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
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
	s := micro.NewService(opts)
	if err := s.Server().Subscribe(
		server.NewSubscriber(
			"go.micro.srv.TickRecorder",
			new(subscriber.Tick),
		),
	); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Println(err)
	}
}
