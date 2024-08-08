package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export all todo items to stdout in specified format",
	RunE:  commandExport,
}

func init() {
	RootCmd.AddCommand(exportCmd)

	exportCmd.Flags().String("format", "json", "Export format. Available formats are: json, csv")
}

func commandExport(cmd *cobra.Command, args []string) error {
	fmt.Println("export called")
	return nil
}
