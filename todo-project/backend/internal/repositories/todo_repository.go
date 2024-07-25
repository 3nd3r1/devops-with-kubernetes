package repositories

import "todo-project-backend/internal/api/models"

type TodoRepository struct {
	todos []models.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: []models.Todo{},
	}
}

func (tr *TodoRepository) GetAll() ([]models.Todo, error) {
	return tr.todos, nil
}

func (tr *TodoRepository) Create(newTodo models.NewTodo) (models.Todo, error) {
	todo := models.Todo{
		Id:        len(tr.todos) + 1,
		Title:     newTodo.Title,
		Completed: false,
	}
	tr.todos = append(tr.todos, todo)

	return todo, nil
}
