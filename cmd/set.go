package cmd

import (
	"fmt"
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/berk-karaal/todo-cli/internal/repository"
	"github.com/spf13/cobra"
	"slices"
	"strconv"
	"strings"
)

var setCmd = &cobra.Command{
	Use:   "set TODO_ID STATUS",
	Short: "Update status of a specific todo",
	Long: `Update status of a specific todo

Available STATUS options are: done, undone, ip`,
	RunE: commandSet,
	Args: cobra.ExactArgs(2),
}

func init() {
	RootCmd.AddCommand(setCmd)
}

func commandSet(cmd *cobra.Command, args []string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	todoId, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("TODO_ID argument must be an integer")
	}

	status := strings.ToLower(args[1])
	if !slices.Contains([]string{"done", "undone", "ip"}, status) {
		return fmt.Errorf("STATUS must be either done, undone or ip")
	}

	statusMap := map[string]repository.TodoStatus{
		"done":   repository.TODO_STATUS_DONE,
		"undone": repository.TODO_STATUS_NOT_STARTED,
		"ip":     repository.TODO_STATUS_IN_PROGRESS,
	}

	todoRepo := repository.TodoRepository{DB: db}
	effectedRowCount, err := todoRepo.UpdateTodoStatus(todoId, statusMap[status])
	if err != nil {
		return err
	}

	fmt.Printf("Updated %d todo\n", effectedRowCount)
	return nil
}
