package publisher

import (
	"fmt"
	"log"

	nats "github.com/nats-io/nats"

	tickproto "github.com/nii236/nii-finance/services/TickRecorder/proto"
)

// PublishTick sends out single tick to telegraf to save to InfluxDB
func PublishTick(t *tickproto.Tick) {
	broker := t.Broker
	last := t.Last
	ask := t.Ask
	bid := t.Bid
	pair := t.Pair
	time := t.Time

	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	msg := fmt.Sprintf("tick,broker=%s,pair=%s ask=%f,bid=%f,last=%f %d", broker, pair, ask, bid, last, time)
	if err := nc.Publish("go.micro.telegraf", []byte(msg)); err != nil {
		log.Println(err)
	}
}
