package src

import (
	"os"

	"log/slog"
)

var Logger *slog.Logger
var logPath string

func init() {
	Logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	// TODO
	logPath = "/ADD/log/PATH"
}
