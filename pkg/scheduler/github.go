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
	//user := ghDetails.User.GetLogin()
	opt := &github.NotificationListOptions{
		All: true,
	}
	notifications, _, err := client.Activity.ListNotifications(ctx, opt)
	if err != nil {
		return err
	}
	for _, notification := range notifications {
		logrus.Debugf("%#v", *notification.Repository.IssuesURL)
		org := notification.Repository.Organization
		repo := notification.Repository
		//logrus.Printf("This is the org: %#v. This is the repo: %#v", org)
		comments, _, err := client.PullRequests.GetComment(ctx, org.GetLogin(), repo.GetName(), 0)
		if err != nil {
			return err
		}
		logrus.Println(comments)
	}
	return nil
}

func setRepoSubTrue(ghDetails *fs.GitHub) {
	// TODO: implement block repos.
	ctx := context.Background()
	client := ghDetails.Client
	//user := ghDetails.User.GetLogin()
	subed := true
	subscription := github.Subscription{
		Subscribed: &subed,
	}

	orgs, _, err := client.Organizations.List(ctx, "", nil)
	if err != nil {
		logrus.Warnf("Failed to list organizations: %s", err)
	}
	for _, org := range orgs {
		repos, _, err := client.Repositories.ListByOrg(ctx, org.GetLogin(), nil)
		if err != nil {
			logrus.Warnf("Failed to get organization repos: %s", err)
		}
		for _, repo := range repos {
			sub, _, err := client.Activity.SetRepositorySubscription(ctx, org.GetLogin(), repo.GetName(), &subscription)
			if err != nil {
				logrus.Warnf("Failed to subscribe to repository: %s (%s)", repo, err)
				continue
			}
			logrus.Debugf("Subscribed to repository: %s", sub.GetRepositoryURL())
		}
	}
}
