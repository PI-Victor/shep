package scheduler

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/PI-Victor/shep/pkg/services"
)

// Scheduler is the general service
type Scheduler struct{}

// NewScheduler returns a new instance of the service structure along with
// defaults.
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// Start starts the bot service.
func (s *Scheduler) Start(cfg *services.Config) error {
	logrus.Info("Starting Shep... ")
	if err := services.NewGitHubClient(cfg); err != nil {
		return err
	}
	go services.SetRepoSubTrue(cfg.GitHub)
	for {
		if err := services.WatchRepos(cfg.GitHub); err != nil {
			return err
		}
		logrus.Debug("Sleeping...")
		time.Sleep(10 * time.Second)
	}
}
