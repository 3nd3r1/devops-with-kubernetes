package main

import (
	"context"
	"os"
	"todo-project-backend/internal/api"
	"todo-project-backend/internal/config"
	"todo-project-backend/internal/database"
	"todo-project-backend/internal/logger"
	"todo-project-backend/internal/nats"
)

func handleError(err error) {
	if err != nil {
		logger.Logger.Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	var err error
	ctx := context.Background()

	err = logger.Init()
	handleError(err)
	err = config.Init(ctx)
	handleError(err)
	err = database.Init()
	handleError(err)
	err = nats.Init()
	handleError(err)

	defer database.Connection.Close()
	defer nats.Connection.Close()

	server, err := api.NewAPI()
	handleError(err)

	err = server.Run()
	handleError(err)
}
