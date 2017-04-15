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
	for {
		err := NewGitHubClient(cfg)
		if err != nil {
			return err
		}
		err = WatchRepos(cfg)
		if err != nil {
			return err
		}
		time.Sleep(10 * time.Second)
	}
}
