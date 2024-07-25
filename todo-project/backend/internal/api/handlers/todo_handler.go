package handlers

import (
	"net/http"
	"todo-project-backend/internal/api/models"
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
	ctx.JSON(http.StatusCreated, createdTodo)
}

func (th *TodoHandler) GetAllTodos(ctx *gin.Context) {
	todos, err := th.TodoRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}
