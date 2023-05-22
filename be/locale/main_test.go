package locale

import (
	"os"
	"testing"

	"github.com/hangnadi/simple-api-project/be/lib/env"
)

func TestMain(m *testing.M) {

	Init(env.Development)
	os.Exit(m.Run())
}
