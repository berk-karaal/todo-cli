/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long:  ``,
	RunE:  commandList,
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("all", "a", false,
		"List all todos including completed ones. (By default lists only incomplete ones)")
	listCmd.Flags().IntP("day", "d", 0,
		"List todos created at date which is N days before today. (default is 0 which means today).\n"+
			"For example 1 means todos created yesterday, 2 means todos created day before yesterday.",
	)
}

func commandList(cmd *cobra.Command, args []string) error {
	fmt.Println("list called")

	return nil
}
