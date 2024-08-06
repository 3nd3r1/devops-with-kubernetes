package api

import (
	"todo-project-backend/internal/api/middleware"
	"todo-project-backend/internal/api/routes"
	"todo-project-backend/internal/config"
	"todo-project-backend/internal/logger"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
}

func NewAPI() (*API, error) {
	router := gin.New()

	router.Use(middleware.RequestLogger())
	router.Use(middleware.ResponseLogger())
	router.Use(gin.Recovery())

	err := routes.SetupRoutes(router)
	if err != nil {
		return nil, err
	}

	return &API{
		Router: router,
	}, nil
}

func (api *API) Run() error {
	logger.Logger.Info("Starting server in port " + config.Config.Port)

	err := api.Router.Run(config.Config.Ip + ":" + config.Config.Port)
	if err != nil {
		return err
	}

	return nil
}
