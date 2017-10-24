package services

import (
	"errors"

	"github.com/ktrysmt/go-bitbucket"
)

// Bitbucket holds the information necessary to create a client connection to
// the bitbucket API.
type Bitbucket struct {
	AppID     string `json:"appID,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`

	IgnoreRepos []string `json:"ignoreRepos,omitempty"`

	Client *bitbucket.Client `json:"-"`
}

// LoadService validates if the service has the necessary config to start
func (b *Bitbucket) LoadService(cfg *Config) error {
	// TODO: some extra validation will be needed here.
	if cfg.Bitbucket == nil || len(cfg.Bitbucket.SecretKey) == 0 {
		return errors.New("bitbucket config not found, skipping init...")
	}
	if err := b.StartService(cfg); err != nil {
		return err
	}
	return nil
}

// StartService starts the bitbucket service
func (b *Bitbucket) StartService(cfg *Config) error {
	b.Client = bitbucket.NewOAuth(cfg.Bitbucket.AppID, cfg.Bitbucket.SecretKey)
	// TODO: this needs an internal error to be surfaced.
	if b.Client == nil {
		return errors.New("An error occured while creating a new client")
	}
	return nil
}
