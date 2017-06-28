package scheduler

import (
	"context"
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
	ctx := context.Background()
	if err := services.NewGitHubClient(ctx, cfg); err != nil {
		return err
	}
	go services.SetRepoSubTrue(ctx, cfg.GitHub)
	for {
		if err := services.WatchRepos(ctx, cfg.GitHub); err != nil {
			return err
		}
		logrus.Debug("Sleeping...")
		time.Sleep(10 * time.Second)
	}
}
