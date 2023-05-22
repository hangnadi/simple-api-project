package repository_test

import (
	"testing"

	. "github.com/simple-api-project/be/internal/repository"
)

func TestSystemLogOpenFile(t *testing.T) {
	systemRepo := NewSystem("")
	systemRepo.LogOpenFile()
}
