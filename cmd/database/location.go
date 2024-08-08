package database

import (
	"fmt"
	"github.com/spf13/cobra"
)

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Prints the absolute path of database being used",
	RunE:  commandLocation,
}

func init() {
	databaseCmd.AddCommand(locationCmd)
}

func commandLocation(cmd *cobra.Command, args []string) error {
	fmt.Println("/path/of/data.sqlite")
	return nil
}
