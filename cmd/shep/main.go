package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/PI-Victor/shep/pkg/cli"
)

var rootCmd = &cobra.Command{
	Use:     "shep",
	Example: "shep - A versatile GitHub bot that runs tests against various CI/CD servers.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(cli.StartCmd)
	rootCmd.Execute()
}

func init() {
	viper.AddConfigPath("$HOME/.shep")
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.json")
	logrus.SetLevel(logrus.DebugLevel)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Failed to read Shep config %s", err)
	}
}
