package routes

import (
	"todo-project-backend/internal/api/handlers"
	"todo-project-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	todoHandler := handlers.NewTodoHandler(repositories.NewTodoRepository())

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/todos", todoHandler.GetAllTodos)
		apiGroup.POST("/todos", todoHandler.CreateTodo)
	}
}
