package zaplogger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	logger.Debug("test", topicField("tp-event"))
}

func unknownField(value string) core.Field {
	return core.Field{
		Type:        core.UnknownType,
		Name:        "unknown",
		ValueString: value,
	}
}

func TestDebugFail(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.NotNil(t, err)
		}
	}()
	logger.Debug("test", unknownField("unknown value"))
}
