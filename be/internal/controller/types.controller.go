package controller

import (
	"github.com/hangnadi/simple-api-project/be/internal/repository"
)

// Service module
type (
	systemSvc struct {
		systemRP repository.System
	}
)

type (
	// A System service provides all functions related for System
	System interface {
		LogOpenFile()
	}
)
