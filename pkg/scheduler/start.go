package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/PI-Victor/shep/pkg/services"
)

var (
	vers       = "v0.1"
	author     = "Victor Palade"
	URL        = "https://github.com/PI-Victor/shep"
	welcomeMsg = fmt.Sprintf(`
 _____ _    _ ______ _____
/ ____| |  | |  ____|  __ \       Automation Bot for VCS systems
| (___| |__| | |__  | |__)|                  %s
\___ \|  __  |  __| |  ___/             by %s
____) | |  | | |____| |          %s
|____/|_|  |_|______|_|
`,
		vers,
		author,
		URL,
	)
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
	logrus.Infof("Starting... \n%s \n", welcomeMsg)
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
