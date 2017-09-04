package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudflavor/shep/pkg/scheduler"
	"github.com/cloudflavor/shep/pkg/services"
)

var (
	defaultCfgDir string
)

// StartCmd starts the bot.
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start Shep",
	Example: `
shep start - Starts the application.
`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			logrus.Fatalf("Failed to read config: %s", err)
		}
		newCfg := services.NewCfg()
		if err := viper.Unmarshal(newCfg); err != nil {
			logrus.Fatalf("An error occured while reading the config: %s", err)
		}
		if err := services.ValidateCfg(newCfg); err != nil {
			logrus.Fatalf("An error occured while validating the config: %s", err)
		}
		newScheduler := scheduler.NewScheduler()
		if err := newScheduler.Start(newCfg); err != nil {
			logrus.Fatalf("an error occured while starting the application: %s", err)
		}
	},
}

// ConfigCmd creates the default configuration of the application.
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "creates a new config with default values.",
	Example: `
Create a default configuration file for the application in the current working
directory.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := services.CreateDefaultCfg(); err != nil {
			logrus.Fatalf("An error occured while generating a default config: %s", err)
		}
	},
}
