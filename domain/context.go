package domain

import (
	"context"
)

var (
	// Ctx is the Browser Context
	Ctx *context.Context
)

// GetContext returns the context
func GetContext() *context.Context {
	return Ctx
}
