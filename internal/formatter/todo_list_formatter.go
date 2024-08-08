package formatter

import (
	"fmt"
	"github.com/berk-karaal/todo-cli/internal/repository"
)

func TodoFormatter(todo repository.Todo) string {
	var status string
	switch todo.Status {
	case repository.TODO_STATUS_NOT_STARTED:
		status = " "
	case repository.TODO_STATUS_IN_PROGRESS:
		status = "IP"
	case repository.TODO_STATUS_DONE:
		status = "X"
	default:
		status = "?"
	}

	return fmt.Sprintf("%d. [%s] %s", todo.Id, status, todo.Name)
}
