package scheduler

import (
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/PI-Victor/shep/pkg/fs"
)

// Scheduler is the general service
type Scheduler struct{}

// NewScheduler returns a new instance of the service structure along with
// defaults.
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// Start starts the bot service.
func (s *Scheduler) Start(cfg *fs.Config) error {
	logrus.Info("Starting Shep... ")

	if err := NewGitHubClient(cfg); err != nil {
		return err
	}
	if err := setRepoSubTrue(cfg.GitHub); err != nil {
		return err
	}
	for {
		if err := WatchRepos(cfg.GitHub); err != nil {
			return err
		}
		logrus.Debug("Sleeping...")
		time.Sleep(10 * time.Second)
	}
}
