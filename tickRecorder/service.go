package main

import (
	"github.com/Sirupsen/logrus"
	micro "github.com/micro/go-micro"
	server "github.com/nii236/nii-forex/tickRecorder/server"
)

var (
	log *logrus.Logger
)

func getOptions(o *micro.Options) {
	o.Server = server.NewNatsServer()
}

func main() {
	log = logrus.New()
	log.Info("Starting tickRecorder service...")
	service := micro.NewService(getOptions)
	service.Run()

}
