package services

import (
	"github.com/ktrysmt/go-bitbucket"
)

// Bitbucket holds the information necessary to create a client connection to
// the bitbucket API.
type Bitbucket struct {
	User     string
	Password string
}

// LoadService validates if the service has the necessary config to start
func (b *Bitbucket) LoadService(cfg *Config) error {
	return nil
}

// StartService starts the bitbucket service
func (b *Bitbucket) StartService(cfg *Config) error {
	return nil
}

// newClient creates a new bitbucket client
func (b *Bitbucket) newClient() error {
	return nil
}

func newBitBucket(cfg *Config) *Bitbucket {
	return nil
}
