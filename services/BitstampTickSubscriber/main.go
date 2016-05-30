package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/micro/go-micro/cmd"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	proto "open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto"

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
	log.Println("Publishing trade data...")
	log.Println(data.(string))

	t := &trade{}
	json.Unmarshal([]byte(data.(string)), t)
	log.Println("Unmarshalled:", t)
	ctx := metadata.NewContext(context.Background(), metadata.MD{"X-User-Id": []string{"BitstampTickSubscriber"}})
	msg := client.NewPublication("go.micro.srv.BitstampRecorder", &proto.Trade{
		Time:   t.Timestamp,
		Price:  t.Price,
		Amount: t.Amount,
	})
	if err := client.Publish(ctx, msg); err != nil {
		log.Println("publish err: ", err)
	}
	log.Println("Done")
}
