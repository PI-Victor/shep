package services

import (
	"encoding/json"
	"errors"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// Service abstracts away all the functionality that a service should implement.
type Service interface {
	LoadService(cfg *Config) error
	StartService(cfg *Config) error
}

// Config holds the configuration options for the application.
type Config struct {
	DebugLevel logrus.Level `json:"debugLevel"`

	Timer int64 `json:"timer"`

	GitHub *GitHub `json:"github"`

	Jenkins *Jenkins `json:"jenkins,omitempty"`

	Bitbucket *Bitbucket `json:"bitbucket,omitempty"`

	Labels []Label `json:"labels,omitempty"`
}

// Label holds information about a label that the bot will add to github to tag
// PRs and issues.
type Label struct {
	Name     string `json:"name,omitempty"`
	Severity int    `json:"severity,omitempty"`
	HexColor string `json:"hexColor,omitempty"`
}

// NewCfg returns a new empty config instance.
func NewCfg() *Config {
	return &Config{}
}

// NewDefaultConfig is the default config that is used to generate a new
// config.json file
func newDefaultCfg() *Config {
	return &Config{
		DebugLevel: logrus.InfoLevel,
		Timer:      45,
		Labels: []Label{
			{
				Name:     "P1",
				Severity: 1,
			}, {
				Name:     "P2",
				Severity: 2,
			}, {
				Name:     "P3",
				Severity: 3,
			}, {
				Name:     "Needs Rebase",
				Severity: 0,
			}, {
				Name:     "Needs Labeling",
				Severity: 0,
			},
		},
	}
}

// ValidateCfg validates the configuration.
func ValidateCfg(cfg *Config) error {
	if cfg.GitHub == nil {
		return errors.New("you need to specify a GitHub token")
	}
	return nil
}

// CreateDefaultCfg creates a default config.json in the current working
// directory.
func CreateDefaultCfg() error {
	cfgDir, err := os.Getwd()
	if err != nil {
		return err
	}

	cfgFile := path.Join(cfgDir, ".shep")
	logrus.Print(cfgFile)
	fh, err := os.Create(cfgFile)
	if err != nil {
		return err
	}
	defer fh.Close()

	newDefaultCfg := newDefaultCfg()
	config, err := json.MarshalIndent(newDefaultCfg, "", " ")
	if err != nil {
		return err
	}
	if _, err := fh.Write(config); err != nil {
		return err
	}

	return nil
}
