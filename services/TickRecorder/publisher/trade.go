package publisher

import (
	"fmt"
	"log"

	nats "github.com/nats-io/nats"

	tickproto "open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto"
)

func PublishTrade(t *tickproto.Trade) {
	broker := t.Broker
	price := t.Price
	time := t.Time
	amount := t.Amount
	pair := t.Pair
	tradeType := "untyped"
	if t.Type == 0 {
		tradeType = "buy"
	} else if t.Type == 1 {
		tradeType = "sell"
	}
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	msg := fmt.Sprintf("trade,broker=%s,type=%s,pair=%s price=%f,amount=%f %d", broker, tradeType, pair, price, amount, time)
	if err := nc.Publish("go.micro.telegraf", []byte(msg)); err != nil {
		log.Println(err)
	}
	log.Println("Published.")
}
