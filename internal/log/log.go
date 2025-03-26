package log

import (
	"log/slog"
	"os"
)

func NewLogger(config Config) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource:   false,
		Level:       nil,
		ReplaceAttr: nil,
	}

	if config.Level == "debug" {
		opts.Level = slog.LevelDebug
	}

	var handler slog.Handler

	switch config.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

func SetGlobalLogger(logger *slog.Logger) {
	slog.SetDefault(logger)
}
