package routes

import (
	"net/http"
	"todo-project-backend/internal/api/handlers"
	"todo-project-backend/internal/database"
	"todo-project-backend/internal/nats"
	"todo-project-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) error {
	todoRepository, err := repositories.NewTodoRepository()
	if err != nil {
		return err
	}
	todoHandler := handlers.NewTodoHandler(todoRepository)

	router.GET("/", func(c *gin.Context) {
		if !database.IsReady || !nats.IsReady {
			c.String(http.StatusServiceUnavailable, "not ready")
			return
		}
		c.String(http.StatusOK, "ok")
	})

	router.GET("/healthz", func(c *gin.Context) {
		if !database.IsReady || !nats.IsReady {
			c.String(http.StatusServiceUnavailable, "not ready")
			return
		}
		c.String(http.StatusOK, "ok")
	})

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/todos", todoHandler.GetAllTodos)
		apiGroup.POST("/todos", todoHandler.CreateTodo)
		apiGroup.PUT("/todos/:id", todoHandler.CompleteTodo)
	}

	return nil
}
