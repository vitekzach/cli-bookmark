package src

import (
	"os"

	"log/slog"
)

var logger *slog.Logger
var logPath string

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	// TODO
	logPath = "/ADD/log/PATH"
}
