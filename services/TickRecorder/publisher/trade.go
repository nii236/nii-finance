package publisher

import (
	"fmt"
	"log"
	"strconv"

	nats "github.com/nats-io/nats"

	tickproto "open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto"
)

func PublishTrade(t *tickproto.Trade) {
	broker := t.Broker
	price := strconv.FormatFloat(t.Price, 'f', -1, 64)
	time := strconv.FormatInt(t.Time, 10)

	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	msg := fmt.Sprintf("trade,broker=%s value=%s %s", broker, price, time)
	if err := nc.Publish("go.micro.telegraf", []byte(msg)); err != nil {
		log.Println(err)
	}
	log.Println("Published.")
}
