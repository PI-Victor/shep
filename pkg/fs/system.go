package fs

import (
	"encoding/json"
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
)

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

// GitHub holds specific information that is used for GitHub integration.
type GitHub struct {
	User  github.User `json:"-"`
	Token string      `json:"token"`

	// A list of URLs that the bot can ignore.
	RepoIgnoreList []string `json:"repoIgnoreList,omitempty"`
	// Holds the client instance details. Internal only.
	Client github.Client `json:"-"`
}

// Redis has information about a redis instance.
type Redis struct{}

// Config holds the configuration options for the application.
type Config struct {
	DebugLevel logrus.Level `json:"debugLevel"`

	IRCServers []IRCSettings `json:"ircServers"`

	GitHub GitHub `json:"gitHub"`

	JenkinsServers []JenkinsSettings `json:"jenkinsServers, omitempty"`

	TravisToken string `json:"travisToken"`

	Redis Redis `json:"redis"`
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
func CreateDefaultCfg(dir string) error {
	var err error
	cfgDir := dir

	if cfgDir == "" {
		cfgDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	cfgFile := path.Join(cfgDir, ".shep.json")
	logrus.Print(cfgFile)
	fh, err := os.Create(cfgFile)
	if err != nil {
		return err
	}
	defer fh.Close()

	newDefaultCfg := newDefaultConfig()
	config, err := json.MarshalIndent(newDefaultCfg, "", " ")
	if err != nil {
		return err
	}
	if _, err := fh.Write(config); err != nil {
		return err
	}

	return nil
}
