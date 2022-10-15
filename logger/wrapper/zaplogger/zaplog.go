package zaplogger

import (
	"fmt"
	"net/url"

	"github.com/w-woong/common/logger/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type rollerWithSink struct {
	// *lumberjack.Logger
	core.Roller
}

func (rollerWithSink) Sync() error {
	return nil
}

func toZapLevel(level core.Level) zapcore.Level {
	var zlevel zapcore.Level
	switch level {
	case core.DebugLevel:
		zlevel = zap.DebugLevel
	case core.InfoLevel:
		zlevel = zap.InfoLevel
	case core.WarnLevel:
		zlevel = zap.WarnLevel
	case core.ErrorLevel:
		zlevel = zap.ErrorLevel
	case core.FatalLevel:
		zlevel = zap.FatalLevel
	case core.PanicLevel:
		zlevel = zap.PanicLevel
	default:
		panic(fmt.Errorf("unknown log level %s", level))
	}
	return zlevel
}

func toZapFields(fields ...core.Field) []zapcore.Field {
	fieldLength := len(fields)
	if fieldLength == 0 {
		var empty []zapcore.Field
		return empty
	}

	var fieldIndex uint8 = 0
	zapFields := make([]zapcore.Field, fieldLength)

	for _, f := range fields {
		switch f.Type {
		case core.StringType:
			zapFields[fieldIndex] = zap.String(f.Name, f.ValueString)
			fieldIndex++
		case core.BytesType:
			zapFields[fieldIndex] = zap.ByteString(f.Name, f.ValueBytes)
			fieldIndex++
		case core.Int32Type:
			zapFields[fieldIndex] = zap.Int32(f.Name, f.ValueInt32)
			fieldIndex++
		case core.Int64Type:
			zapFields[fieldIndex] = zap.Int64(f.Name, f.ValueInt64)
			fieldIndex++
		default:
			panic(fmt.Errorf("unknown field type, %d(%s)", f.Type, f.Name))
		}
	}

	return zapFields[:fieldIndex]
}

func NewLogger(level core.Level, stdOut bool, fileName string, roller core.Roller) *ZapLogger {

	var zlevel zapcore.Level = toZapLevel(level)
	var err error

	config := zap.NewProductionConfig()
	config.Sampling = nil
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.CallerKey = "caller"
	// encoderConfig.FunctionKey = "funcName"
	encoderConfig.StacktraceKey = ""

	config.Level = zap.NewAtomicLevelAt(zlevel)
	config.EncoderConfig = encoderConfig
	config.Encoding = "json"

	if roller != nil {
		zap.RegisterSink("roller", func(*url.URL) (zap.Sink, error) {
			// if l, ok := roller.(*lumberjack.Logger); ok {
			// 	return rollerWithSink{
			// 		Logger: l,
			// 	}, nil
			// }
			// return nil, errors.New("unknown sinker")
			return rollerWithSink{
				Roller: roller,
			}, nil
		})

		config.OutputPaths = []string{fmt.Sprintf("roller:%s", fileName)}
		if stdOut {
			config.OutputPaths = append(config.OutputPaths, "stderr")
		}
	}

	_globallogger, err := config.Build(zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(_globallogger)

	return &ZapLogger{
		logger: _globallogger,
	}
}

type ZapLogger struct {
	logger *zap.Logger
}

func (z *ZapLogger) Close() {
	if z.logger != nil {
		z.logger.Sync()
	}
}

func (z *ZapLogger) Debug(message string, fields ...core.Field) {
	z.logger.Debug(message, toZapFields(fields...)...)
}

func (z *ZapLogger) Info(message string, fields ...core.Field) {
	z.logger.Info(message, toZapFields(fields...)...)
}

func (z *ZapLogger) Warn(message string, fields ...core.Field) {
	z.logger.Warn(message, toZapFields(fields...)...)
}

func (z *ZapLogger) Error(message string, fields ...core.Field) {
	z.logger.Error(message, toZapFields(fields...)...)
}

func (z *ZapLogger) Fatal(message string, fields ...core.Field) {
	z.logger.Fatal(message, toZapFields(fields...)...)
}

func (z *ZapLogger) Panic(message string, fields ...core.Field) {
	z.logger.Panic(message, toZapFields(fields...)...)
}
