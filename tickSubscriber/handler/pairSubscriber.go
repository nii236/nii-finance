package handler

import (
	"github.com/Sirupsen/logrus"
	"github.com/nii236/nii-finance/tickSubscriber/proto"
	"golang.org/x/net/context"
)

type Tick struct{}

var log *logrus.Logger

func (t *Tick) StartSubscription(ctx context.Context, req *tickSubscriber.SubscribeRequest, res *tickSubscriber.SubscribeResponse) error {
	log = logrus.New()
	log.Println("Received tick subscribe request for pair:", req.Pair)

	return nil
}
