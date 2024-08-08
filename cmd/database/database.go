package database

import (
	"github.com/berk-karaal/todo-cli/cmd"
	"github.com/spf13/cobra"
)

// databaseCmd represents the database command
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database related commands",
}

func init() {
	cmd.RootCmd.AddCommand(databaseCmd)
}
