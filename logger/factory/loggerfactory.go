package factory

import (
	"fmt"

	"github.com/w-woong/common/logger/core"
	"github.com/w-woong/common/logger/wrapper/logruslogger"
	"github.com/w-woong/common/logger/wrapper/lumberlogger"
	"github.com/w-woong/common/logger/wrapper/zaplogger"
)

type Factory interface {
	CreateLogger(loggerType core.LoggerType,
		level string, stdOut bool,
		fileName string, maxSize int, maxBackup int, maxAge int,
		compress bool) (core.Logger, error)
}

func CreateLogger(loggerType core.LoggerType,
	level core.Level, stdOut bool,
	fileName string, roller core.Roller) (core.Logger, error) {

	switch loggerType {
	case core.ZapLogger:
		return zaplogger.NewLogger(level, stdOut, fileName, roller), nil
	case core.LogrusLogger:
		return logruslogger.NewLogger(level, stdOut, true, fileName, roller), nil
	default:
		return nil, fmt.Errorf("unknown logger type, %d", loggerType)
	}
}

func CreateRoller(fileName string, maxSize int, maxBackup int, maxAge int, compress bool) core.Roller {
	if len(fileName) == 0 {
		return nil
	}
	return lumberlogger.NewLogger(fileName, maxSize, maxBackup, maxAge, compress)
}
