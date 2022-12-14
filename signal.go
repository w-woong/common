package common

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/w-woong/common/logger"
)

// StartSignalStopper starts and wait for signals.
// Once received a signal, that included in input signals, then calls stop method of stopper.
func StartSignalStopper(stopper Stopper, signals ...os.Signal) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, signals...)
	go func() {
		sig := <-sigs
		logger.Info(fmt.Sprintf("signal %v", sig))
		stopper.Stop()
	}()
}

// StartTicker calls tick on interval set by ticker.
// Close done channel to terminate.
func StartTicker(done chan bool, ticker *time.Ticker, tick func(time.Time)) {
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				tick(t)
			}
		}
	}()
}
