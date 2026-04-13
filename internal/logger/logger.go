package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func NewLogger(env string) *slog.Logger {
	var logger *slog.Logger
	var options *slog.HandlerOptions

	switch env {
	case envLocal:
		options = &slog.HandlerOptions{
			Level: slog.LevelDebug,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					t := a.Value.Time()
					return slog.String(slog.TimeKey, t.Format("04:05"))
				}
				return a
			},
		}

		logger = slog.New(slog.NewTextHandler(os.Stdout, options))
	case envDev:
		options = &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}

		logger = slog.New(slog.NewJSONHandler(os.Stdout, options))
	case envProd:
		options = &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}

		logger = slog.New(slog.NewJSONHandler(os.Stdout, options))
	}

	return logger
}
