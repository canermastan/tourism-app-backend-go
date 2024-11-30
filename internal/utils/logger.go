package utils

import (
	"context"
)

type Logger interface {
	Info(ctx context.Context, message string)
	Error(ctx context.Context, message string)
}
