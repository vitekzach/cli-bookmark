package src

import (
	"os"

	"log/slog"
)

var logger *slog.Logger
var logPath string

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	// TODO log into file instead
	// TODO put real log path here
	logPath = "/ADD/log/PATH"
}
