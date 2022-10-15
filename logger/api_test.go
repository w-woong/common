package logger_test

import (
	"sync"
	"testing"

	"github.com/w-woong/common/logger"
)

func TestDebug(t *testing.T) {
	logger.Debug("hello")
}

func TestDebugKafka(t *testing.T) {
	numRoutines := 100
	numLoggings := 1000

	wg := sync.WaitGroup{}

	wg.Add(numRoutines)
	for k := 0; k < numRoutines; k++ {
		go func(wg *sync.WaitGroup) {
			for i := 0; i < numLoggings; i++ {
				logger.Debug("hello")
			}
			wg.Done()
		}(&wg)
	}

	wg.Add(numRoutines)
	for k := 0; k < numRoutines; k++ {
		go func(wg *sync.WaitGroup) {
			for i := 0; i < numLoggings; i++ {
				logger.Debug("hello",
					logger.TopicField("tp-event"),
					logger.ValueField([]byte("this is a message")),
					logger.ClientIpField("123.123.123.123"),
				)
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}
