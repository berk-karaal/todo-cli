package cmd

import (
	"fmt"
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/berk-karaal/todo-cli/internal/repository"
	"github.com/spf13/cobra"
	"strconv"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename TODO_ID NEW_NAME",
	Short: "Rename a todo item",
	RunE:  commandRename,
	Args:  cobra.ExactArgs(2),
}

func init() {
	RootCmd.AddCommand(renameCmd)
}

func commandRename(cmd *cobra.Command, args []string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	todoId, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("TODO_ID argument must be an integer")
	}

	newName := args[1]

	todoRepo := repository.TodoRepository{DB: db}
	effectedRowCount, err := todoRepo.UpdateTodoName(todoId, newName)
	if err != nil {
		return err
	}

	fmt.Printf("Updated %d todo\n", effectedRowCount)
	return nil
}
