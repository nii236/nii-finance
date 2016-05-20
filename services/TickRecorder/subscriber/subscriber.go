package subscriber

import (
	"log"

	proto "gitlab.com/open-algot/open-algot-platform/services/TickRecorder/proto"
	"golang.org/x/net/context"
)

type Tick struct{}

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
