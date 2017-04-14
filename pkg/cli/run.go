package cli

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/PI-Victor/shep/pkg/fs"
	"github.com/PI-Victor/shep/pkg/service"
)

var (
	cfgDir string
)

// StartCmd starts the bot.
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start Shep",
	Example: `
shep start - Starts the bot with the configuration provided in
~/.shep/config.json.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			logrus.Fatalf("Failed to read config %s. See `shep config`.", err)
		}
		newCfg := fs.NewConfig()
		if err := viper.Unmarshal(newCfg); err != nil {
			logrus.Fatalf("An error occured while reading the config: %s", err)
		}
		newService := service.NewService()
		newService.Start(newCfg)
	},
}

// ConfigCmd creates the default configuration of the application.
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "creates a new config with default values.",
	Example: `
Create a default configuration file for the application. If you omit --dir, the
configuration is created in the current working directory of the application.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := fs.CreateDefaultCfg(cfgDir); err != nil {
			logrus.Fatalf("an error occured while generating a default config: %s", err)
		}
	},
}

func init() {
	ConfigCmd.PersistentFlags().StringVar(&cfgDir, "dir", "", "set the dir where the default config should be created.")
}
