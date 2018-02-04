package services

import (
	"context"
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

// Validate validates if the service has the necessary config to start
func (b *Bitbucket) Validate(ctx context.Context) error {

	if err := b.startService(); err != nil {
		return err
	}
	return nil
}

// StartService starts the bitbucket service
func (b *Bitbucket) startService() error {
	// TODO: some extra validation will be needed here.
	if b.Client == nil || len(b.SecretKey) == 0 {
		return errors.New("bitbucket config not found, skipping init")
	}
	// TODO: this needs an internal error to be surfaced.
	b.Client = bitbucket.NewOAuth(b.AppID, b.SecretKey)

	if b.Client == nil {
		return errors.New("an error occured while creating a new client")
	}
	return nil
}

func (b *Bitbucket) Run(ctx context.Context) error {
	return nil
}

func NewBitbucketService(cfg *Config) *Bitbucket {
	return &Bitbucket{
		AppID:       cfg.Bitbucket.AppID,
		SecretKey:   cfg.Bitbucket.SecretKey,
		IgnoreRepos: cfg.Bitbucket.IgnoreRepos,
	}
}
