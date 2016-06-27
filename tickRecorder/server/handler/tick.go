package handler

import (
	"time"

	"github.com/Sirupsen/logrus"
	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/nii236/nii-finance/tickRecorder/proto"
	"golang.org/x/net/context"
)

type tickData struct {
	log    *logrus.Logger
	influx influx.Client
}

func (t *tickData) TickHandler(ctx context.Context, tick *tickRecorder.Tick) error {
	var err error
	t.log.Infoln("Received data")
	tags := map[string]string{"pair": "AUDUSD"}
	fields := map[string]interface{}{
		"bid":  tick.Bid,
		"ask":  tick.Ask,
		"last": tick.Last,
	}

	point, err := influx.NewPoint("tick_data", tags, fields, time.Unix(0, tick.Time))

	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  "tick",
		Precision: "ns",
	})

	bp.AddPoint(point)
	t.log.Infoln("Created batch point:", bp)
	if influxErr := t.influx.Write(bp); influxErr != nil {
		t.log.Error(influxErr)
	}

	return err
}

func NewTickHandler(log *logrus.Logger, client influx.Client) *tickData {
	return &tickData{log: log, influx: client}
}
