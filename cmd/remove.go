package cmd

import (
	"fmt"
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/berk-karaal/todo-cli/internal/repository"
	"github.com/spf13/cobra"
	"strconv"
)

var removeCmd = &cobra.Command{
	Use:   "remove TODO_ID",
	Short: "Remove specific todo from database",
	RunE:  commandRemove,
	Args:  cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(removeCmd)
}

func commandRemove(cmd *cobra.Command, args []string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	todoId, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("TODO_ID argument must be an integer")
	}

	todoRepo := repository.TodoRepository{DB: db}
	effectedRowCount, err := todoRepo.DeleteTodoById(todoId)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted %d todo\n", effectedRowCount)
	return nil
}
