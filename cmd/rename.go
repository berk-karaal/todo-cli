package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename TODO_ID NEW_NAME",
	Short: "Rename a todo item",
	RunE:  commandRename,
	Args:  cobra.ExactArgs(2),
}

func init() {
	RootCmd.AddCommand(renameCmd)
}

func commandRename(cmd *cobra.Command, args []string) error {
	fmt.Println("rename called")
	return nil
}
