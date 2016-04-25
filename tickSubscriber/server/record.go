package server

import (
	"os"
	"os/signal"

	"github.com/gofinance/ib"
)

var (
	e *ib.Engine
)

func exec() {
	options := &ib.EngineOptions{
		Client:           0,
		Gateway:          "192.168.99.100:4003",
		DumpConversation: false,
	}

	var err error
	e, err = ib.NewEngine(*options)

	defer e.Stop()

	if err != nil {
		log.Infoln("ERROR!", err)

		return
	}

	defer e.Stop()

	usdjpy := &ib.Contract{
		Symbol:       "USD",
		SecurityType: "CASH",
		Exchange:     "IDEALPRO",
		Currency:     "JPY",
	}

	im, err := ib.NewInstrumentManager(e, *usdjpy)

	if err != nil {
		log.Infoln(err)
	}

	go last(im)

	// go timeChan(ctm)

	for {
	}

}

func spinup() {

	log.Infoln("Starting IB app")

	signal.Notify(sigChan, os.Interrupt)

	log.Infoln("Launching Engine")
	go processor()

ControlLoop:
	for {
		select {
		case <-sigChan:
			log.Error("OS interrupt received")
			close(shutdown)
			sigChan = nil

		case err := <-complete:
			log.Infof("IB Engine Completed: Error[%s]", err)
			break ControlLoop
		}
	}
	log.Println("IB app complete")

}

func last(i *ib.InstrumentManager) {

	for {
		select {
		case <-i.Refresh():
			log.Infoln(i.Bid())
			log.Infoln(i.Ask())
			log.Infoln(i.Last())
		}
	}

}

func timeChan(ctm *ib.CurrentTimeManager) {
	for {
		select {
		case <-ctm.Refresh():
			log.Infoln("CTM Refresh")
			log.Infoln(ctm.Time())
		}
	}
}
