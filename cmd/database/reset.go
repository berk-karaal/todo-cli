package database

import (
	"fmt"
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
	fmt.Println("reset called")
	return nil
}
