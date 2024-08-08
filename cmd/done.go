package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done TODO_ID",
	Short: "Mark a todo item as done",
	RunE:  commandDone,
	Args:  cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(doneCmd)
}

func commandDone(cmd *cobra.Command, args []string) error {
	fmt.Println("Marking command as COMPLETED " + args[0])
	return nil
}
