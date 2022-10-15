package logruslogger_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/w-woong/common/logger/core"
	"github.com/w-woong/common/logger/wrapper/logruslogger"
)

var (
	logger *logruslogger.Logger
)

func setup() error {
	logger = logruslogger.NewLogger(core.DebugLevel, true, true, "", nil)
	return nil
}

func shutdown() {
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Println(err)
		shutdown()
		os.Exit(1)
	}

	exitCode := m.Run()

	shutdown()
	os.Exit(exitCode)
}
