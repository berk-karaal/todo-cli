package formatter

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"github.com/berk-karaal/todo-cli/internal/repository"
	"strconv"
	"time"
)

type TodoExport struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

// parseToTodoExport is used to prepare data objects before export in various formats in order to keep uniform
// export values.
func parseToTodoExport(todos []repository.Todo) []TodoExport {
	var exportData []TodoExport
	for _, todo := range todos {
		exportData = append(exportData, TodoExport{
			Id:        todo.Id,
			Name:      todo.Name,
			Status:    string(todo.Status),
			CreatedAt: todo.CreatedAt.Format(time.RFC3339),
		})
	}
	return exportData
}

// JsonExportFormatter returns json string to export given repository.Todo objects
func JsonExportFormatter(todos []repository.Todo) (string, error) {
	exportData := parseToTodoExport(todos)

	jsonBytes, err := json.Marshal(exportData)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

// CsvExportFormatter returns csv string to export given repository.Todo objects
func CsvExportFormatter(todos []repository.Todo) (string, error) {
	exportData := parseToTodoExport(todos)

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"id", "name", "status", "createdAt"}
	err := writer.Write(header)
	if err != nil {
		return "", err
	}
	for _, v := range exportData {
		err = writer.Write([]string{
			strconv.Itoa(v.Id), v.Name, v.Status, v.CreatedAt,
		})
		if err != nil {
			return "", err
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		return "", err
	}

	return buf.String(), nil
}
