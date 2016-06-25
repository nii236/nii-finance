package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/micro/go-micro/cmd"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"

	proto "github.com/nii236/nii-finance/services/TickRecorder/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"github.com/micro/go-micro/client"
	pusher "github.com/pusher-community/pusher-websocket-go"
)

//Example data: {"price": 469.0, "timestamp": "1464340140", "amount": 1.80891143, "type": 0, "id": 11226014}
type trade struct {
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
	Amount    float64 `json:"amount"`
	Type      int     `json:"type"`
	ID        int     `json:"id"`
	Broker    string
}

func main() {
	cmd.Init()
	wg := &sync.WaitGroup{}
	p := pusher.New("de504dc5763aeef9ff52")

	wg.Add(1)
	ticker := p.Subscribe("live_trades")
	ticker.Bind("trade", publish)
	wg.Wait()

}

func publish(data interface{}) {
	log.Println("Publishing trade data...", data.(string))

	t := &trade{}
	json.Unmarshal([]byte(data.(string)), t)
	t.Broker = "bitstamp"
	ctx := metadata.NewContext(context.Background(), metadata.MD{"X-User-Id": []string{"BitstampTickSubscriber"}})
	now := time.Now().UnixNano()
	msg := client.NewPublication("go.micro.srv.TickRecorder.Trade", &proto.Trade{
		Time:   now,
		Price:  t.Price,
		Amount: t.Amount,
		Broker: t.Broker,
		Type:   int32(t.Type),
		Pair:   "BTCUSD",
	})
	if err := client.Publish(ctx, msg); err != nil {
		log.Println("publish err: ", err)
	}
}
