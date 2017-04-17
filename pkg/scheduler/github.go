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
	logrus.Println(&cfg.GitHub.Client)
	logrus.Println(cfg.GitHub.User.GetBio())
	client := &cfg.GitHub.Client
	orgs, _, err := client.Organizations.List(ctx, "", nil)
	if err != nil {
		return err
	}
	for _, org := range orgs {
		projects, _, err := cfg.GitHub.Client.Organizations.ListProjects(ctx, org.GetName(), nil)
		if err != nil {
			return err
		}

		for _, proj := range projects {
			logrus.Println(proj)
			//logrus.Printf("This is the Org: %s and repo: %s", *org.Name, *proj.Name)
		}
	}
	return nil
}
