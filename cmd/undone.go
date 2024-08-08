package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone TODO_ID",
	Short: "Mark a todo items as undone",
	RunE:  commandUndone,
	Args:  cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(undoneCmd)
}

func commandUndone(cmd *cobra.Command, args []string) error {
	fmt.Println("undone called")
	return nil
}
