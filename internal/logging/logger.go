package logging

import (
	"context"
	"log"
	"os"
)

type lKey struct{}

var (
	logContextKey = lKey{}
	defaultLogger = newDefaultLogger()
)

// DefaultLogger exports the default logger instance defined here.
func DefaultLogger() *log.Logger {
	return defaultLogger
}

func newDefaultLogger() *log.Logger {
	return log.New(os.Stdout, "rapi: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// NewContext embeds a logger into the context.
func NewContext(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, logContextKey, logger)
}

// FromContext retrieves a logger from the context or returns the default..
func FromContext(ctx context.Context) *log.Logger {
	if logger, ok := ctx.Value(logContextKey).(*log.Logger); ok {
		return logger
	}
	return defaultLogger
}
