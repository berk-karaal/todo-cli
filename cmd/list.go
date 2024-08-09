package cmd

import (
	"fmt"
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/berk-karaal/todo-cli/internal/formatter"
	"github.com/berk-karaal/todo-cli/internal/repository"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long:  ``,
	RunE:  commandList,
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().IntP("day", "d", 0,
		"List todos created at date which is N days before today. (default is 0 which means today).\n"+
			"For example 1 means todos created yesterday, 2 means todos created day before yesterday.",
	)
}

func commandList(cmd *cobra.Command, args []string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	todoRepo := repository.NewTodoRepository(db)

	dayFlag, err := cmd.Flags().GetInt("day")
	if err != nil {
		return err
	}

	now := time.Now()
	minTime := time.Date(
		now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local,
	).AddDate(0, 0, -dayFlag)

	todos, err := todoRepo.ListTodosByCreatedAt(minTime, minTime.AddDate(0, 0, 1))
	if err != nil {
		return err
	}

	for _, todo := range todos {
		fmt.Println(formatter.TodoFormatter(todo))
	}

	return nil
}
