package repository

import (
	"database/sql"
	"fmt"
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

// ListTodosByCreatedAt returns Todos from database. Given `min` and `max` parameters are used to filter Todos
// by their creation time.
func (r *TodoRepository) ListTodosByCreatedAt(min, max time.Time) ([]Todo, error) {
	if min.IsZero() || max.IsZero() {
		return nil, fmt.Errorf("min or max Time parameters must be non-zero value")
	}

	rows, err := r.DB.Query(
		"SELECT id, name, status, createdAt FROM todos WHERE createdAt > ? AND createdAt < ?",
		min.Unix(), max.Unix(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var (
			id        int
			name      string
			status    string
			createdAt int
		)
		err := rows.Scan(&id, &name, &status, &createdAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, Todo{
			Id:        id,
			Name:      name,
			Status:    status,
			CreatedAt: time.Unix(int64(createdAt), 0),
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
