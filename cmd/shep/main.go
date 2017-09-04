package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudflavor/shep/pkg/cli"
)

const asciiChar = `
 _____ _    _ ______ _____
/ ____| |  | |  ____|  __ \
| (___| |__| | |__  | |__)|
\___ \|  __  |  __| |  ___/
____) | |  | | |____| |
|____/|_|  |_|______|_|
`

var (
	rootCmd = &cobra.Command{
		Use: "shep",
		Example: fmt.Sprintf(`
%s
shep - A versatile automation bot for VCS systems that runs tests against various CI/CD servers.
`, asciiChar),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

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
