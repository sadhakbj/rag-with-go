package logger

import (
	"log/slog"
	"os"
	"time"
)

// InitializeLogger initializes a logger with configurable log level and timestamp.
func InitializeLogger(name string, addSource bool, level slog.Level) *slog.Logger {
	handler := &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
				a.Value = slog.StringValue(time.Now().Format(time.RFC3339)) // More readable timestamp
			}
			return a
		},
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, handler).WithAttrs([]slog.Attr{
		slog.String("service", name),
	}))
	return logger
}
