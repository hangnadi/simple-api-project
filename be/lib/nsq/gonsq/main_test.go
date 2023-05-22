package gonsq_test

import (
	"os"
	"testing"

	. "github.com/simple-api-project/be/lib/nsq/gonsq"
)

func TestMain(m *testing.M) {

	TFuncPatch()

	os.Exit(m.Run())
}
