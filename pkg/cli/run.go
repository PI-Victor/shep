package cli

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/PI-Victor/shep/pkg/fs"
	"github.com/PI-Victor/shep/pkg/service"
)

// StartCmd starts the bot.
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start Shep",
	Example: `shep start - Starts the bot with the configuration provided in
~/.shep/config.json.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config := fs.NewConfig()
		if err := viper.Unmarshal(config); err != nil {
			logrus.Fatalf("An error occured while reading the config: %s", err)
		}
		service.Start(config)
	},
}

// ConfigCmd creates the default configuration of the application.
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "config",
	Example: `Use shep config --dir to create a default configuration file for
the application.
  `,
}
