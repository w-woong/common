package core

type Logger interface {
	Debug(message string, fields ...Field)
	Info(message string, fields ...Field)
	Warn(message string, fields ...Field)
	Error(message string, fields ...Field)
	Fatal(message string, fields ...Field)
	Panic(message string, fields ...Field)
}

type Closer interface {
	Close()
}

type LoggerType uint8

const (
	ZapLogger LoggerType = iota
	LogrusLogger
)

type Roller interface {
	Write(p []byte) (n int, err error)
	Close() error
	Rotate() error
}
