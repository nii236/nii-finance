package subscriber

import (
	"log"

	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	proto "github.com/nii236/nii-finance/services/TickRecorder/proto"
	"github.com/nii236/nii-finance/services/TickRecorder/publisher"
)

// Tick is a struct that contains Tick handlers
type Tick struct {
	Client client.Client
}

// Handle will respond to relevant messages on the topic it is registered
func (e *Tick) Handle(ctx context.Context, msg *proto.Tick) error {
	log.Print("Handler received tick data. Publishing...")
	publisher.PublishTick(msg)
	return nil
}
