package main

import (
	"github.com/Sirupsen/logrus"
	micro "github.com/micro/go-micro"
	server "github.com/nii236/nii-forex/tickSubscriber/server"
)

var (
	natsURL = "192.168.99.100:32773"
	log     *logrus.Logger
)

func getOptions(o *micro.Options) {
	o.Server = server.NewNatsServer()
}

func main() {
	log = logrus.New()
	service := micro.NewService(getOptions)
	service.Run()
}
