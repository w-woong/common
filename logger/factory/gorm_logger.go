package factory

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/w-woong/common/logger/core"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	gormlogger.Config
	logger core.Logger
}

func NewGormLogger(level core.Level, logger core.Logger) *GormLogger {
	gormLogLevel := gormlogger.Silent
	switch level {
	case core.DebugLevel:
		gormLogLevel = gormlogger.Info
	case core.InfoLevel:
		gormLogLevel = gormlogger.Info
	case core.WarnLevel:
		gormLogLevel = gormlogger.Warn
	case core.ErrorLevel:
		gormLogLevel = gormlogger.Error
	case core.FatalLevel:
		gormLogLevel = gormlogger.Error
	case core.PanicLevel:
		gormLogLevel = gormlogger.Error
	}
	return &GormLogger{
		Config: gormlogger.Config{LogLevel: gormLogLevel},
		logger: logger,
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.logger.Info(fmt.Sprint(append([]interface{}{msg}, data...)...))
	}
}

func (l GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.logger.Warn(fmt.Sprint(append([]interface{}{msg}, data...)...))
	}
}

func (l GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.logger.Error(fmt.Sprint(append([]interface{}{msg}, data...)...))
	}
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		l.logger.Error(err.Error(), core.WithStringField("sql", sql), core.WithInt64Field("rows", rows))
		// if rows == -1 {
		// 	logger.Error(err.Error(), core.WithStringField("sql", sql), core.WithInt64Field("rows", rows))
		// 	// l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		// } else {
		// 	logger.Error(err.Error(), core.WithStringField("sql", sql), core.WithInt64Field("rows", rows))
		// 	// l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		// }
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		l.logger.Warn(slowLog, core.WithStringField("sql", sql), core.WithInt64Field("rows", rows))
	case l.LogLevel == gormlogger.Info:
		sql, rows := fc()
		l.logger.Info("SQL", core.WithStringField("sql", sql), core.WithInt64Field("rows", rows))
	}
}
