package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use: "clean",
	Short: "Clean old todo items which are created before than the specified days ago. " +
		"You can use is to clean your database from unused data.",
	RunE: commandClean,
}

func init() {
	RootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().IntP("day", "d", 30,
		"Day limit for task to be cleaned if they are created before than the given days ago.")
}

func commandClean(cmd *cobra.Command, args []string) error {
	fmt.Println("clean called")
	return nil
}
