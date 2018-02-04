package scheduler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/cloudflavor/shep/pkg/services"
)

var (
	vers       = "v0.1"
	author     = "Cloudflavor Org"
	url        = "https://github.com/cloudflavor/shep"
	welcomeMsg = fmt.Sprintf(`
 _____ _    _ ______ _____
/ ____| |  | |  ____|  __ \       Automation Bot for SCM systems
| (___| |__| | |__  | |__)|                  %s
\___ \|  __  |  __| |  ___/             by %s
____) | |  | | |____| |          %s
|____/|_|  |_|______|_|
`,
		vers,
		author,
		url,
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
	var (
		ticker = time.Ticker{}
		ctx    = context.Background()
		c      = make(chan os.Signal, 1)
	)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			ticker.Stop()
			logrus.Infof("Received %s, exiting.", sig.String())
			os.Exit(0)
		}
	}()

	scmServices, err := loadServices(ctx, cfg)
	if err != nil {
		return err
	}

	if cfg.Timer < 45 {
		return fmt.Errorf("interval must be greater than 45s, got %ds", cfg.Timer)
	}
	t := (time.Duration(cfg.Timer) * time.Second)
	duration := time.NewTicker(t)

	for range duration.C {
		for _, service := range scmServices {
			go service.Run(ctx)
		}
		logrus.Debugf("Sleeping... %s", time.Now())
	}
	return nil
}

func loadServices(ctx context.Context, cfg *services.Config) ([]services.Service, error) {
	scmServices := []services.Service{}
	// TODO: move this to github service validation instead of here.
	if cfg.GitHub != nil && cfg.GitHub.Token != "" {
		// TODO: abstract away the listeners (github, gitlab, etc).

		if err := services.NewGitHubClient(ctx, cfg); err != nil {
			return nil, err
		}
		// TODO: move this away from service start as well.
		githubService := services.NewGithubService(cfg)
		go services.SetRepoSubTrue(ctx, cfg.GitHub)

		scmServices = append(scmServices, githubService)
	}

	if cfg.Bitbucket != nil && cfg.Bitbucket.SecretKey != "" {
		bitbucketService := services.NewBitbucketService(cfg)
		scmServices = append(scmServices, bitbucketService)
	}
	if len(scmServices) == 0 {
		return nil, errors.New("no SCM services defined")
	}
	return scmServices, nil
}
