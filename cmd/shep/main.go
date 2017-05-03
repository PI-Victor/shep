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
	rootCmd.AddCommand(
		cli.StartCmd,
		cli.ConfigCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	viper.SetConfigFile(".shep")
	viper.SetConfigType("json")
	logrus.SetLevel(logrus.DebugLevel)
}
