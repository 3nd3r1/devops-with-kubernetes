package api

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"todo-project-backend/internal/api/routes"
	"todo-project-backend/internal/config"
)

type API struct {
	Config *config.Config
	Logger *slog.Logger
	Router *gin.Engine
}

func NewAPI(config *config.Config, logger *slog.Logger) *API {
	router := gin.Default()
	routes.SetupRoutes(router)

	return &API{
		Config: config,
		Logger: logger,
		Router: router,
	}
}

func (api *API) Run() error {
	api.Logger.Info("Server started in port " + api.Config.Port)

	err := api.Router.Run(api.Config.Ip + ":" + api.Config.Port)
	if err != nil {
		return err
	}

	return nil
}
