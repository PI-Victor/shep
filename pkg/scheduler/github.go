package scheduler

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/PI-Victor/shep/pkg/fs"
)

var (
	lastCheck time.Time
)

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
	// NOTE: i don't think i should use context.Background like this, need to look
	// into it.
	ctx := context.Background()
	client := ghDetails.Client
	opt := &github.NotificationListOptions{All: true}

	notifications, _, err := client.Activity.ListNotifications(ctx, opt)
	if err != nil {
		return err
	}
	for _, notification := range notifications {
		parts := strings.Split(*notification.Subject.URL, "/")
		repo := notification.Repository.GetName()
		last := parts[len(parts)-1]
		org := parts[len(parts)-4]
		id, err := strconv.Atoi(last)
		id = int(id)
		if err != nil {
			return err
		}
		comments, _, err := client.Issues.ListComments(ctx, org, repo, id, nil)
		if err != nil {
			logrus.Warning(err)
			continue
		}
		for _, comment := range comments {
			if lastCheck.Before(comment.GetCreatedAt()) {
				if err := checkComment(comment.GetBody(), org, repo, id); err != nil {
					logrus.Warningf("failed to apply action %s", err)
				}
			}
		}
	}
	lastCheck = time.Now()
	return nil
}

func checkComment(body string, org, repo string, id int) error {
	match, err := regexp.Match("[[a-zA-Z]+]", []byte(body))
	if err != nil {
		return err
	}
	if match {
		comm := strings.Trim(body, "[]")
		if comm == "test" {
			commentPR(org, repo, id)
		}
	}
	return nil
}

func commentPR(org, repo string, id int) {
	logrus.Info("test")
}

func mergeCommit() error {
	return nil
}

func setRepoSubTrue(ghDetails *fs.GitHub) {
	// TODO: implement black-lister repos.
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
	// refactor this to avoid O(n)
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
