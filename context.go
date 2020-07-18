package logging

import "context"

type contextKey string

const (
	loggerKey = contextKey("log")
)

// FromContext will return the logger for the specified context or the default one
func FromContext(ctx context.Context) *Log {
	log := ctx.Value(loggerKey)
	if log != nil {
		return log.(*Log)
	}

	return DefaultLog
}

// NewContext creates a child context of the supplied context embedding the *Log. This *Log can be retrieved with the FromContext
func NewContext(ctx context.Context, log *Log) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}
