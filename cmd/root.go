package cmd

import (
	"errors"
	"fmt"
	"github.com/adrg/xdg"
	"github.com/berk-karaal/todo-cli/internal/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Simple todo cli app for daily tasks",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return setup()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// setup function makes necessary setup operations for app to run.
func setup() error {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	configDir := path.Join(userConfigDir, "todo-cli")
	dataDir := path.Join(xdg.DataHome, "todo-cli")

	// Create parent directories of config file if not exist
	err = os.MkdirAll(configDir, 0766)
	if err != nil {
		return err
	}

	// Create  data directory and its parent directories if not exists
	err = os.MkdirAll(dataDir, 0766)
	if err != nil {
		return err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.SetDefault("database.location", path.Join(dataDir, "todo.sqlite"))

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// create a new config file with default settings
			err = viper.WriteConfigAs(path.Join(configDir, "config.yaml"))
			if err != nil {
				return fmt.Errorf("unable to write config file: %w", err)
			}
		}
	}

	// Make sure database tables are created
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	err = database.CreateTables(db)
	if err != nil {
		return err
	}
	db.Close()

	return nil
}
