package logruslogger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/w-woong/common/logger/core"
)

func toLogrusFields(fields ...core.Field) logrus.Fields {
	// if len(fields) == 0 {
	// 	return nil
	// }
	logrusFields := make(logrus.Fields)
	for _, f := range fields {
		switch f.Type {
		case core.StringType:
			logrusFields[f.Name] = f.ValueString
		case core.BytesType:
			logrusFields[f.Name] = f.ValueBytes
		case core.Int32Type:
			logrusFields[f.Name] = f.ValueInt32
		case core.Int64Type:
			logrusFields[f.Name] = f.ValueInt64
		default:
			panic(fmt.Errorf("unknown field type, %d(%s)", f.Type, f.Name))
		}
	}

	if _, fileName, line, ok := runtime.Caller(2); ok {
		// fnName := RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
		fileName = filepath.Base(fileName)
		logrusFields["caller"] = fmt.Sprintf("%s:%d", fileName, line)
		// "caller":"zapwrapper/zaplog_test.go:19"
	}

	return logrusFields
}

func NewLogger(level core.Level, stdOut bool, isJason bool,
	fileName string, roller core.Roller) *Logger {

	logger := logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	if isJason {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// log.SetOutput(os.Stdout)
	if stdOut {
		if roller != nil {
			mw := io.MultiWriter(os.Stderr, roller)
			logger.SetOutput(mw)
		} else {
			logger.SetOutput(os.Stderr)
		}
	} else {
		if roller != nil {
			logger.SetOutput(roller)
		}
	}

	// Only log the warning severity or above.
	parsedLevel, err := logrus.ParseLevel(string(level))
	if err != nil {
		panic(fmt.Errorf("unknown log level %s", string(level)))
	}
	logger.SetLevel(parsedLevel)
	// logger.AddHook(ContextHook{})
	logger.SetReportCaller(false)

	return &Logger{
		logger: logger,
	}
}

type Logger struct {
	logger *logrus.Logger
}

func (z *Logger) Debug(message string, fields ...core.Field) {
	// if len(fields) == 0 {
	// 	z.logger.Debug(message)
	// 	return
	// }
	z.logger.WithFields(toLogrusFields(fields...)).Debug(message)
}

func (z *Logger) Info(message string, fields ...core.Field) {
	// if len(fields) == 0 {
	// 	z.logger.Info(message)
	// 	return
	// }
	z.logger.WithFields(toLogrusFields(fields...)).Info(message)
}

func (z *Logger) Warn(message string, fields ...core.Field) {
	// if len(fields) == 0 {
	// 	z.logger.Warn(message)
	// 	return
	// }
	z.logger.WithFields(toLogrusFields(fields...)).Warn(message)
}

func (z *Logger) Error(message string, fields ...core.Field) {
	// if len(fields) == 0 {
	// 	z.logger.Error(message)
	// 	return
	// }
	z.logger.WithFields(toLogrusFields(fields...)).Error(message)
}

func (z *Logger) Fatal(message string, fields ...core.Field) {
	// if len(fields) == 0 {
	// 	z.logger.Fatal(message)
	// 	return
	// }
	z.logger.WithFields(toLogrusFields(fields...)).Fatal(message)
}

func (z *Logger) Panic(message string, fields ...core.Field) {
	// if len(fields) == 0 {
	// 	z.logger.Panic(message)
	// 	return
	// }
	z.logger.WithFields(toLogrusFields(fields...)).Panic(message)
}

var RE_stripFnPreamble = regexp.MustCompile(`^.*\.(.*)$`)

// func LogrusFields() log.Fields {
// 	fnName := "<unknown>"
// 	fileName := "<unknown>"
// 	// fileName := "<unknown>"
// 	// Skip this function, and fetch the PC and file for its parent
// 	pc, fileName, line, ok := runtime.Caller(1)
// 	if ok {
// 		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
// 		fileName = filepath.Base(fileName)
// 	}

// 	f := log.Fields{
// 		"fileName": fmt.Sprintf(fileName+"(%v)", line),
// 		"funcName": fnName,
// 	}
// 	return f
// }

// func LogrusHTTPFields(url string, reqBody, resBody []byte) log.Fields {
// 	fnName := "<unknown>"
// 	fileName := "<unknown>"
// 	lURL := ""
// 	var lReqBody string
// 	var lResBody string
// 	// fileName := "<unknown>"
// 	// Skip this function, and fetch the PC and file for its parent
// 	pc, fileName, line, ok := runtime.Caller(1)
// 	if ok {
// 		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
// 		fileName = filepath.Base(fileName)
// 	}

// 	if len(url) > 0 {
// 		lURL = url
// 	}
// 	if len(reqBody) > 0 {
// 		lReqBody = string(reqBody)
// 	}
// 	if len(resBody) > 0 {
// 		lResBody = string(resBody)
// 	}

// 	f := log.Fields{
// 		"fileName": fmt.Sprintf(fileName+"(%v)", line),
// 		"funcName": fnName,
// 		"url":      lURL,
// 		"reqBody":  lReqBody,
// 		"resBody":  lResBody,
// 	}
// 	return f
// }

// func LogrusHTTPFieldsV2(url string, ip string, reqBody, resBody []byte) log.Fields {
// 	fnName := "<unknown>"
// 	fileName := "<unknown>"
// 	lURL := ""
// 	var lReqBody string
// 	var lResBody string
// 	// fileName := "<unknown>"
// 	// Skip this function, and fetch the PC and file for its parent
// 	pc, fileName, line, ok := runtime.Caller(1)
// 	if ok {
// 		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
// 		fileName = filepath.Base(fileName)
// 	}

// 	if len(url) > 0 {
// 		lURL = url
// 	}
// 	if len(reqBody) > 0 {
// 		lReqBody = string(reqBody)
// 	}
// 	if len(resBody) > 0 {
// 		lResBody = string(resBody)
// 	}

// 	f := log.Fields{
// 		"fileName": fmt.Sprintf(fileName+"(%v)", line),
// 		"funcName": fnName,
// 		"url":      lURL,
// 		"clientIP": ip,
// 		"reqBody":  lReqBody,
// 		"resBody":  lResBody,
// 	}
// 	return f
// }

// func LogrusHTTPFieldsPOS(url string, ip, shopCod, posNum, requestDump string, reqBody, resBody []byte) log.Fields {
// 	fnName := "<unknown>"
// 	fileName := "<unknown>"
// 	lURL := ""
// 	var lReqBody string
// 	var lResBody string
// 	// fileName := "<unknown>"
// 	// Skip this function, and fetch the PC and file for its parent
// 	pc, fileName, line, ok := runtime.Caller(1)
// 	if ok {
// 		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
// 		fileName = filepath.Base(fileName)
// 	}

// 	if len(url) > 0 {
// 		lURL = url
// 	}
// 	if len(reqBody) > 0 {
// 		lReqBody = string(reqBody)
// 	}
// 	if len(resBody) > 0 {
// 		lResBody = string(resBody)
// 	}

// 	f := log.Fields{
// 		"file_name":     fmt.Sprintf(fileName+"(%v)", line),
// 		"func_name":     fnName,
// 		"url":           lURL,
// 		"client_ip":     ip,
// 		"shop_cod":      shopCod,
// 		"pos_num":       posNum,
// 		"reqquest_body": lReqBody,
// 		"response_body": lResBody,
// 		"request_dump":  requestDump,
// 	}
// 	return f
// }

// func LogrusKafkaFields(failure, topic string, partition int32, offset int64, value []byte) log.Fields {
// 	fnName := "<unknown>"
// 	fileName := "<unknown>"

// 	// Skip this function, and fetch the PC and file for its parent
// 	pc, fileName, line, ok := runtime.Caller(1)
// 	if ok {
// 		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
// 		fileName = filepath.Base(fileName)
// 	}

// 	f := log.Fields{
// 		"fileName":  fmt.Sprintf(fileName+"(%v)", line),
// 		"funcName":  fnName,
// 		"kafka_msg": failure,
// 		"topic":     topic,
// 		"partition": partition,
// 		"offset":    offset,
// 		"value":     string(value),
// 	}
// 	return f
// }
