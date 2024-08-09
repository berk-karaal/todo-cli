package cmd

import (
	"errors"
	"fmt"
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
		return initViper()
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

func init() {
}

func initViper() error {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	configDir := path.Join(userConfigDir, "todo-cli")

	err = os.MkdirAll(configDir, 0766)
	if err != nil {
		return err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.SetDefault("database.location", "./todo.sqlite")

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
	return nil
}
