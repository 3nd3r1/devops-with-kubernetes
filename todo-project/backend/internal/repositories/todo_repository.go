package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"todo-project-backend/internal/api/models"
	"todo-project-backend/internal/config"
	"todo-project-backend/internal/logger"

	_ "github.com/lib/pq"
)

type TodoRepository struct {
	database *sql.DB
}

func NewTodoRepository() (*TodoRepository, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Config.PostgresUser, config.Config.PostgresPassword,
		config.Config.PostgresHost, config.Config.PostgresPort,
		config.Config.PostgresDatabase))

	if err != nil {
		return nil, err
	}

	for {
		err = db.Ping()
		if err == nil {
			break
		}
		logger.Logger.Info("Failed connecting to database, retrying in 10 seconds...")
		time.Sleep(10 * time.Second)
	}

	logger.Logger.Info("Connected to database")
	logger.Logger.Info("Running migrations")

	_, err = db.Query("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, title TEXT, completed BOOLEAN)")
	if err != nil {
		return nil, err
	}

	logger.Logger.Info("Migrations complete")

	return &TodoRepository{
		database: db,
	}, nil
}

func (tr *TodoRepository) GetAll() ([]models.Todo, error) {
	todos := []models.Todo{}

	rows, err := tr.database.Query("SELECT id, title, completed FROM todos")
	if err == sql.ErrNoRows {
		return todos, nil
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Completed)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	err = rows.Err()
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (tr *TodoRepository) Create(newTodo models.NewTodo) (models.Todo, error) {
	var addedTodo models.Todo
	err := tr.database.QueryRow("INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, title, completed",
		newTodo.Title, false).Scan(&addedTodo.Id, &addedTodo.Title, &addedTodo.Completed)

	return addedTodo, err
}
