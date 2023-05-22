package server_test

import (
	"testing"

	. "github.com/hangnadi/simple-api-project/be/lib/server"
	"github.com/stretchr/testify/assert"
)

func TestGetIPAddress(t *testing.T) {
	assert.NotNil(t, GetIPAddress())
}
