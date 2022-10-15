package logger_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/w-woong/common/logger"
)

func setup() error {
	logger.Open("debug", false, "./agent.log", 100, 10, 10, false)
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
