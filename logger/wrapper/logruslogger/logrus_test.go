package logruslogger_test

import (
	"testing"

	"github.com/w-woong/common/logger/core"
)

func topicField(value string) core.Field {
	return core.Field{
		Type:        core.StringType,
		Name:        "topic",
		ValueString: value,
	}
}
func TestDebug(t *testing.T) {
	logger.Debug("hello")
	logger.Debug("hello with fields", topicField("topic"))
}
