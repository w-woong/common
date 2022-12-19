package utils

import (
	"regexp"
	"runtime"
)

// Regex to extract just the function name (and not the module path)
var RE_stripFnPreamble = regexp.MustCompile(`^.*\.(.*)$`)

// Trace Functions
func GetFunctionName() string {
	return GetFunctionNameWithSkip(1)
}

// Trace Functions
func GetFunctionNameWithSkip(skip int) string {
	fnName := "<unknown>"
	pc, _, _, ok := runtime.Caller(skip)
	if ok {
		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
	}

	return fnName
}
