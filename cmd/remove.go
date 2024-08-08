package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove TODO_ID",
	Short: "Remove specific todo item from database",
	RunE:  commandRemove,
	Args:  cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(removeCmd)
}

func commandRemove(cmd *cobra.Command, args []string) error {
	fmt.Println("remove called")
	return nil
}
