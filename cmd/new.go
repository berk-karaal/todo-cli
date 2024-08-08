package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
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
	fmt.Println("Creating new todo with name " + args[0])
	return nil
}
