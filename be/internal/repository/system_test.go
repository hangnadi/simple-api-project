package repository_test

import (
	"testing"

	. "github.com/hangnadi/simple-api-project/be/internal/repository"
)

func TestSystemLogOpenFile(t *testing.T) {
	systemRepo := NewSystem("")
	systemRepo.LogOpenFile()
}
