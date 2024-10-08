package handlers

import (
	"encoding/json"
	"net/http"

	"todo-project-backend/internal/api/models"
	"todo-project-backend/internal/database"
	"todo-project-backend/internal/logger"
	"todo-project-backend/internal/nats"
	"todo-project-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	TodoRepository *repositories.TodoRepository
}

func NewTodoHandler(tr *repositories.TodoRepository) *TodoHandler {
	return &TodoHandler{
		TodoRepository: tr,
	}
}

func (th *TodoHandler) CreateTodo(ctx *gin.Context) {
	var newTodo models.NewTodo

	if !database.IsReady {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "not ready"})
		return
	}

	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(newTodo.Title) > 140 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is over 140 characters"})
		return
	}

	createdTodo, err := th.TodoRepository.Create(newTodo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdTodoJson, err := json.Marshal(createdTodo)
	if err != nil {
		logger.Logger.Error("Error marshalling created todo: " + err.Error())
	} else {
		err = nats.Connection.Publish("todos", createdTodoJson)
		if err != nil {
			logger.Logger.Error("Error publishing created todo to nats: " + err.Error())
		} else {
			logger.Logger.Info("Published created todo to nats")
		}
	}

	ctx.JSON(http.StatusCreated, createdTodo)
}


func (th *TodoHandler) CompleteTodo(ctx *gin.Context) {
	if !database.IsReady {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "not ready"})
		return
	}

	id := ctx.Param("id")

	completedTodo, err := th.TodoRepository.Complete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	completedTodoJson, err := json.Marshal(completedTodo)
	if err != nil {
		logger.Logger.Error("Error marshalling completed todo: " + err.Error())
	} else {
		err = nats.Connection.Publish("todos", completedTodoJson)
		if err != nil {
			logger.Logger.Error("Error publishing completed todo to nats: " + err.Error())
		} else {
			logger.Logger.Info("Published completed todo to nats")
		}
	}

	ctx.JSON(http.StatusOK, completedTodo)
}

func (th *TodoHandler) GetAllTodos(ctx *gin.Context) {
	if !database.IsReady {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "not ready"})
		return
	}

	todos, err := th.TodoRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}
