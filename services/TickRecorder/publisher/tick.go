package publisher

import (
	"fmt"
	"log"
	"strconv"

	nats "github.com/nats-io/nats"

	tickproto "open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto"
)

func PublishTick(t *tickproto.Tick) {
	broker := t.Broker
	last := strconv.FormatFloat(t.Last, 'f', -1, 64)
	time := strconv.FormatInt(t.Time, 10)

	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	msg := fmt.Sprintf("tick,broker=%s value=%s %s", broker, last, time)
	if err := nc.Publish("go.micro.telegraf", []byte(msg)); err != nil {
		log.Println(err)
	}
	log.Println("Published.")
}
