package lumberlogger

import "github.com/natefinch/lumberjack"

func NewLogger(fileName string, maxSize int, maxBackup int, maxAge int, compress bool) *lumberjack.Logger {

	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackup,
		MaxAge:     maxAge,   //days
		Compress:   compress, // disabled by default
	}
}
