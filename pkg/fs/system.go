package fs

import "github.com/Sirupsen/logrus"

// Config holds the configuration options for the application.
type Config struct {
	DebugLevel logrus.Level `json:"debugLevel"`

	GitHubUser  string `json:"githubUser"`
	GitHubToken string `json:"gitHubToken"`

	JenkinsURL   string `json:"jenkinsURL"`
	JenkinsUser  string `json:"jenkinsUser"`
	JenkinsToken string `json:"jenkinsToken"`

	IRCUser   string `json:"ircUser"`
	IRCServer string `json:"ircServer"`

	TravisToken string `json:"travisToken"`

	// URL List of repositories to watch.
	WatchList []string `json:"watchList"`
}

// NewConfig returns a new empty config instance.
func NewConfig() *Config {
	return &Config{}
}

// NewDefaultConfig is the default config that is used to generate a new
// config.json file
func NewDefaultConfig() *Config {
	return &Config{
		DebugLevel: logrus.InfoLevel,
	}
}
