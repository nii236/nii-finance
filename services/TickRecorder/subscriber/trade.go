package subscriber

import (
	"log"

	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	proto "open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto"
	"open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/publisher"
)

// Trade is a struct that contains Trade handlers
type Trade struct {
	c client.Client
}

// Handle will respond to relevant messages on the topic it is registered
func (e *Trade) Handle(ctx context.Context, msg *proto.Trade) error {
	log.Print("TickRecorder received trade data. Publishing...")
	publisher.PublishTrade(msg)
	return nil
}
