package scheduler

import (
	"context"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/PI-Victor/shep/pkg/fs"
)

var lastCheck time.Time

// NewGitHubClient returns a valid instance of a new GitHub client.
func NewGitHubClient(cfg *fs.Config) error {
	ctx := context.Background()
	tokenSrc := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: cfg.GitHub.Token,
		},
	)
	tc := oauth2.NewClient(ctx, tokenSrc)
	client := github.NewClient(tc)
	cfg.GitHub.Client = client
	authClient, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return err
	}

	cfg.GitHub.User = *authClient
	return nil
}

// WatchRepos watches the repositories in the organization that the bot is part
// of for events such as comments or PRs.
func WatchRepos(ghDetails *fs.GitHub) error {
	ctx := context.Background()
	client := ghDetails.Client
	opt := &github.NotificationListOptions{
		All: true,
	}
	notifications, _, err := client.Activity.ListNotifications(ctx, opt)
	if err != nil {
		return err
	}
	for _, n := range notifications {
		logrus.Debug(n)
	}
	return nil
}

func setRepoWatchTrue(ghDetails *fs.GitHub) error {
	ctx := context.Background()
	client := ghDetails.Client
	user := ghDetails.User
	subscription := github.Subscription{}
	// TODO: list all repos available (only org for now) and set subscription to max
	sub, err := client.Activity.SetRepositorySubscription(ctx, user.GetLogin(), subscription)
	if err != nil {
		return err
	}
	logrus.Debug(client, sub)
	return nil
}
