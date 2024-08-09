package database

import (
	"fmt"
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/berk-karaal/todo-cli/internal/repository"
	"github.com/spf13/cobra"
	"time"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use: "clean",
	Short: "Clean old todo items which are created before than the specified days ago. " +
		"You can use is to clean your database from unused data.",
	RunE: commandClean,
}

func init() {
	databaseCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().IntP("day", "d", 30,
		"Day limit for task to be cleaned if they are created before than the given days ago.")
}

func commandClean(cmd *cobra.Command, args []string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	dayFlag, err := cmd.Flags().GetInt("day")
	if err != nil {
		return err
	}

	now := time.Now()
	createdAtLimit := time.Date(
		now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local,
	).AddDate(0, 0, -dayFlag)

	todoRepo := repository.NewTodoRepository(db)
	deletedTodoCount, err := todoRepo.DeleteTodoByCreatedAtSmallerThan(createdAtLimit)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted %d todos created before %s.\n", deletedTodoCount, createdAtLimit.Format(time.DateOnly))
	return nil
}
