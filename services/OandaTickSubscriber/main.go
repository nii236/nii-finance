package main

import (
	"log"
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	proto "gitlab.com/open-algot/open-algot-platform/services/TickRecorder/proto"
	"golang.org/x/net/context"
)

type pairslice []string

var pairs pairslice

func main() {
	cmd.Init()
	wg := sync.WaitGroup{}
	wg.Add(1)
	t := time.NewTicker(1 * time.Second)
	for range t.C {
		log.Println("Publishing mock tick data...")
		ctx := metadata.NewContext(context.Background(), map[string]string{
			"X-User-Id": "john",
			"X-From-Id": "script",
		})
		msg := client.NewPublication("go.micro.srv.TickRecorder", &proto.Tick{
			Time:   123,
			Bid:    1,
			Ask:    2,
			Last:   3,
			Pair:   "AUDUSD",
			Broker: "Oanda",
		})
		if err := client.Publish(ctx, msg); err != nil {
			log.Println("pub err: ", err)
		}
		log.Println("Done")
	}
	wg.Wait()
}
