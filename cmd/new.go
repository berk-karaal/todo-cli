package cmd

import (
	"fmt"
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/berk-karaal/todo-cli/internal/repository"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new TODO_TEXT",
	Short: "Create a new todo",
	RunE:  commandNew,
	Args:  cobra.RangeArgs(1, 1),
}

func init() {
	RootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolP("completed", "c", false, "Create todo as completed")
}

func commandNew(cmd *cobra.Command, args []string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	todoRepo := repository.NewTodoRepository(db)
	todo, err := todoRepo.CreateTodo(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Added '%s'\n", todo.Name)
	return nil
}
