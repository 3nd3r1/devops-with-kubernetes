package main

import (
	"context"
	"log/slog"
	"os"
	"todo-project-backend/internal/api"
	"todo-project-backend/internal/config"
)

func main() {
	ctx := context.Background()

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	logger := slog.New(jsonHandler)

	config := config.NewConfig()
	if err := config.ParseEnv(ctx); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	server := api.NewAPI(config, logger)

	if err := server.Run(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
