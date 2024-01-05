package src

import (
	"os"

	"log/slog"
)

var Logger *slog.Logger

func init() {
	Logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
}
