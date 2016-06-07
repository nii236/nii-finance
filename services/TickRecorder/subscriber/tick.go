package subscriber

import (
	"log"

	"golang.org/x/net/context"
	proto "open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto"
)

// Tick is a struct that contains Tick handlers
type Tick struct{}

// Handle will respond to relevant messages on the topic it is registered
func (e *Tick) Handle(ctx context.Context, msg *proto.Tick) error {
	log.Print("Handler received tick")
	log.Println("Time", msg.Time)
	log.Println("Bid", msg.Bid)
	log.Println("Ask", msg.Ask)
	log.Println("Last", msg.Last)
	log.Println("Pair", msg.Pair)
	log.Println("Broker", msg.Broker)
	return nil
}
