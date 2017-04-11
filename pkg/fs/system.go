package fs

var (
	defaultCfg = `
{
  "debugLevel": 3,
  "githubUser": "",
  "githubToken": "",
  "jenkinsURL": "",
  "jenkinsAPI": "",
  "jenkinsToken": "",
  "travisToken": "",
  "watchList": [],
}
`
)

// Config holds the configuration options for the application.
type Config struct {
	DebugLevel int `json:"debugLevel"`

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

// NewConfig returns a new config instance.
func NewConfig() *Config {
	return &Config{}
}
