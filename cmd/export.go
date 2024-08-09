package cmd

import (
	"fmt"
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/berk-karaal/todo-cli/internal/formatter"
	"github.com/berk-karaal/todo-cli/internal/repository"
	"github.com/spf13/cobra"
)

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
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	todoRepo := repository.NewTodoRepository(db)
	todos, err := todoRepo.ListAllTodos()
	if err != nil {
		return err
	}

	exportFormat, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}

	var output string
	var outputErr error
	switch exportFormat {
	case "json":
		output, outputErr = formatter.JsonExportFormatter(todos)
	case "csv":
		output, outputErr = formatter.CsvExportFormatter(todos)
	default:
		return fmt.Errorf("invalid export format")
	}
	if outputErr != nil {
		return outputErr
	}

	fmt.Println(output)
	return nil
}
