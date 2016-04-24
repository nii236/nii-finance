package subscriber

import (
	"log"

	"github.com/nii236/nii-forex/tickRecorder/proto"

	"golang.org/x/net/context"
)

type Tick struct{}

func (e *Tick) Handle(ctx context.Context, msg *tick.Message) error {
	log.Print("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *tick.Message) error {
	log.Print("Function Received message: ", msg.Say)
	return nil
}
