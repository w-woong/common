package zaplogger_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/w-woong/common/logger/core"
	"github.com/w-woong/common/logger/wrapper/zaplogger"
)

var (
	logger *zaplogger.ZapLogger
)

func setup() error {
	logger = zaplogger.NewLogger(core.DebugLevel, true, "", nil)
	return nil
}

func shutdown() {
	logger.Close()
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
