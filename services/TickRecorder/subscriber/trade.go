package subscriber

import (
	"log"

	"golang.org/x/net/context"
	proto "open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto"
)

type Trade struct{}

func (e *Trade) Handle(ctx context.Context, msg *proto.Trade) error {
	log.Print("TickRecorder received trade data")
	log.Println("Time:	", msg.Time)
	log.Println("Price:	", msg.Price)
	log.Println("Amount:	", msg.Amount)
	log.Println("Type:	", msg.Type)
	return nil
}
