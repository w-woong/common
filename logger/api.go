package logger

import (
	"github.com/w-woong/common/logger/core"
	"github.com/w-woong/common/logger/factory"
)

var _logger core.Logger

func init() {
	_logger, _ = factory.CreateLogger(core.ZapLogger, core.DebugLevel, true, "", nil)
}

func Open(level string, stdOut bool,
	fileName string, maxSize int, maxBackup int, maxAge int, compress bool) {

	var err error

	roller := factory.CreateRoller(fileName, maxSize, maxBackup, maxAge, compress)

	_logger, err = factory.CreateLogger(core.ZapLogger, core.Level(level), stdOut, fileName, roller)
	if err != nil {
		panic(err)
	}
}

func Close() {
	if _logger == nil {
		return
	}
	if closer, ok := _logger.(core.Closer); ok {
		closer.Close()
	}
}

func Logger() core.Logger {
	return _logger
}

func OpenGormLogger(level string) *factory.GormLogger {
	if level == "" {
		level = string(core.ErrorLevel)
	}
	return factory.NewGormLogger(core.Level(level), _logger)
}

func Debug(message string, fields ...core.Field) {
	_logger.Debug(message, fields...)
}
func Info(message string, fields ...core.Field) {
	_logger.Info(message, fields...)
}
func Warn(message string, fields ...core.Field) {
	_logger.Warn(message, fields...)
}
func Error(message string, fields ...core.Field) {
	_logger.Error(message, fields...)
}
func Fatal(message string, fields ...core.Field) {
	_logger.Fatal(message, fields...)
}
func Panic(message string, fields ...core.Field) {
	_logger.Panic(message, fields...)
}

// Field
//

func TopicField(value string) core.Field {
	return core.WithStringField("topic", value)
}

func ValueField(value []byte) core.Field {
	return core.WithBytesField("value", value)
}

func PartitionField(value int32) core.Field {
	return core.WithInt32Field("partition", value)
}

func OffsetField(value int64) core.Field {
	return core.WithInt64Field("offset", value)
}

func UrlField(value string) core.Field {
	return core.WithStringField("url", value)
}
func ClientIpField(value string) core.Field {
	return core.WithStringField("client_ip", value)
}
func ReqBodyField(value []byte) core.Field {
	return core.WithBytesField("req_body", value)
}
func ResBodyField(value []byte) core.Field {
	return core.WithBytesField("res_body", value)
}
func ShopCodeField(value string) core.Field {
	return core.WithStringField("shop_cod", value)
}
func PosNumberField(value string) core.Field {
	return core.WithStringField("pos_num", value)
}
