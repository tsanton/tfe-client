package tfe_test

import (
	"log"
	"os"
	"testing"

	logt "github.com/sirupsen/logrus/hooks/test"

	u "github.com/tsanton/tfe-client/tfe/utilities"
)

var (
	logger u.ILogger
)

func TestMain(m *testing.M) {
	log.Println("Main test setup")

	logger, _ = logt.NewNullLogger()

	log.Println("Running tests")
	exitVal := m.Run()
	log.Println("Main test teardown")

	os.Exit(exitVal)
}
