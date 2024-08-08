package repository

import (
	"database/sql"
	"time"
)

type TodoRepository struct {
	DB *sql.DB
}

type Todo struct {
	Id        int
	Name      string
	Status    string
	CreatedAt time.Time
}

type TodoStatus string

const (
	TODO_STATUS_NOT_STARTED = "N"
	TODO_STATUS_IN_PROGRESS = "IP"
	TODO_STATUS_DONE        = "D"
)

func NewTodoRepository(db *sql.DB) TodoRepository {
	return TodoRepository{DB: db}
}

func (r *TodoRepository) CreateTodo(name string) (Todo, error) {
	todoCreatedAt := time.Now()
	result, err := r.DB.Exec(
		"INSERT INTO todos (name, status, createdAt) VALUES (?,?,?)",
		name, TODO_STATUS_NOT_STARTED, todoCreatedAt.Unix())
	if err != nil {
		return Todo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Todo{}, err
	}

	return Todo{
		Id:        int(id),
		Name:      name,
		Status:    TODO_STATUS_NOT_STARTED,
		CreatedAt: todoCreatedAt,
	}, nil
}
