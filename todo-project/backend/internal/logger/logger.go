package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func Init() error {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	Logger = slog.New(jsonHandler)

	return nil
}
