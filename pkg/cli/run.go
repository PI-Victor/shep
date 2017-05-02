package cli

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/PI-Victor/shep/pkg/fs"
	"github.com/PI-Victor/shep/pkg/scheduler"
)

var (
	cfgDir        string
	defaultCfgDir string
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
		// only search the dir passed as parameter so that we don't load some
		// already created config by mistake.
		if cfgDir != "" {
			viper.AddConfigPath(cfgDir)
		} else {
			viper.AddConfigPath("$HOME/shep")
			viper.AddConfigPath("/etc/shep")
			viper.AddConfigPath(".")
		}
		if err := viper.ReadInConfig(); err != nil {
			logrus.Fatalf("Failed to read config: %s", err)
		}
		newCfg := fs.NewCfg()
		if err := viper.Unmarshal(newCfg); err != nil {
			logrus.Fatalf("An error occured while reading the config: %s", err)
		}
		if err := fs.ValidateCfg(newCfg); err != nil {
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
Create a default configuration file for the application. If you omit --dir, the
configuration is created in the current working directory of the application.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := fs.CreateDefaultCfg(cfgDir); err != nil {
			logrus.Fatalf("An error occured while generating a default config: %s", err)
		}
	},
}

func init() {
	StartCmd.PersistentFlags().StringVar(&cfgDir, "config", "", "specify the config dir.")
	ConfigCmd.PersistentFlags().StringVar(&defaultCfgDir, "dir", "", "set the dir where the default config should be created.")
}
