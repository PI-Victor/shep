package fs

import "github.com/Sirupsen/logrus"

// IRCSettings holds information about one or more IRC servers that the bot
// could join.
type IRCSettings struct {
	IRCServers []IRCServer `json:"ircServers"`
}

// IRCServer holds information regarding one or more IRC servers that the bot
// should connect to.
type IRCServer struct {
	ServerName string   `json:"serverName"`
	Nick       string   `json:"nick"`
	Channels   []string `json:"channels"`
}

// JenkinsSettings holds the information about one or more Jenkins servers that
// the bot should send or retrieve information from.
type JenkinsSettings struct {
	JenkinsURL   string `json:"jenkinsURL"`
	JenkinsUser  string `json:"jenkinsUser"`
	JenkinsToken string `json:"jenkinsToken"`
}

// Config holds the configuration options for the application.
type Config struct {
	DebugLevel logrus.Level `json:"debugLevel"`

	IRCServers []IRCSettings `json:"ircServers"`

	JenkinsServers []JenkinsSettings `json:"jenkinsServers"`

	GitHubUser  string `json:"githubUser"`
	GitHubToken string `json:"gitHubToken"`
	// A list of URLs that the bot can ignore.
	GitHubIgnoreList []string `json:"gitHubIgnoreList"`

	TravisToken string `json:"travisToken"`
}

// NewConfig returns a new empty config instance.
func NewConfig() *Config {
	return &Config{}
}

// NewDefaultConfig is the default config that is used to generate a new
// config.json file
func newDefaultConfig() *Config {
	return &Config{
		DebugLevel: logrus.InfoLevel,
	}
}

// CreateDefaultCfg creates a default config.json in the current working
// directory.
func CreateDefaultCfg() error {
	return nil
}
