package logger_test

import (
	"testing"

	"github.com/w-woong/common/logger"
)

func BenchmarkDebug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.Debug("hello")
	}
}

func BenchmarkDebug_WithOption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.Debug("hello with fields", logger.TopicField("tp-event"))
		// log.Debug("hello", logger.Topic("tp-event"), logger.Value([]byte("tp-value")))
	}
}

func BenchmarkDebug_Parallel(b *testing.B) {
	var names []string = []string{"1", "2"}
	for _, name := range names {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				logger.Debug(name)
			}
		})
	}
}
