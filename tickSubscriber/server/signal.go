package server

import "errors"

// processor handles shutdown signals
func processor() error {

	log.Infoln("Processor - Starting")

	task()

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Infoln("Processor - Completed")
	return nil

}

// Task begins fetching from URL url
func task() error {
	// Defer the send on the channel so it happens
	// regardless of how this function terminates.
	var err error
	defer func() {

		// Capture any potential panic.
		if r := recover(); r != nil {
			log.Infoln("Processor - Panic", r)
		}

		// Signal the goroutine we have shutdown.
		complete <- err
	}()

	go exec()

	for {
		if checkShutdown() {
			return errors.New("Early Shutdown")
		}
	}
}

// checkShutdown checks the shutdown flag to determine
// if we have been asked to interrupt processing.
func checkShutdown() bool {
	select {
	case <-shutdown:

		// We have been asked to shutdown cleanly.
		e.Stop()
		return true

	default:

		// If the shutdown channel was not closed,
		// presume with normal processing.
		return false
	}
}
