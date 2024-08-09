package database

import (
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Factory reset to the database",
	RunE:  commandReset,
}

func init() {
	databaseCmd.AddCommand(resetCmd)
}

func commandReset(cmd *cobra.Command, args []string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = database.DropTables(db)
	if err != nil {
		return err
	}

	err = database.CreateTables(db)
	if err != nil {
		return err
	}
	return nil
}
