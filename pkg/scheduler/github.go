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
func WatchRepos(cfg *fs.Config) error {
	ctx := context.Background()
	client := cfg.GitHub.Client
	orgs, _, err := client.Organizations.List(ctx, cfg.GitHub.User.GetLogin(), nil)
	if err != nil {
		return err
	}
	for _, org := range orgs {
		repos, _, err := client.Repositories.ListByOrg(ctx, org.GetLogin(), nil)
		if err != nil {
			return err
		}
		for _, repo := range repos {
			_, _, err := client.PullRequests.GetComment(ctx, org.GetLogin(), "shep_test", 0)
			if err != nil {
				return err
			}
			comments, _, err := client.PullRequests.ListComments(ctx, org.GetLogin(), repo.GetName(), 0, nil)
			if err != nil {
				return err
			}
			logrus.Debugf("the org: %s, repo: %s, comments: %#v", org.GetLogin(), repo.GetName(), comments)
			for _, com := range comments {
				logrus.Debugf("Comment: %#v", com.GetBody())
			}
		}
	}
	return nil
}
