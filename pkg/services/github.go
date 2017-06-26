package services

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var (
	lastCheck time.Time
)

// GitHub holds specific information that is used for GitHub integration.
type GitHub struct {
	User  github.User `json:"-"`
	Token string      `json:"token"`

	// A list of URLs that the bot can ignore.
	IgnoreRepos []string `json:"ignoreRepos,omitempty"`
	// Holds the client instance details. Internal only.
	Client *github.Client `json:"-"`
}

// NewGitHubClient returns a valid instance of a new GitHub client.
func NewGitHubClient(cfg *Config) error {
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
func WatchRepos(ghDetails *GitHub) error {
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
		// TODO: replace this with a proper group regexp.
		parts := strings.Split(*notification.Subject.URL, "/")
		repo := notification.Repository.GetName()
		last := parts[len(parts)-1]
		org := parts[len(parts)-4]
		id, err := strconv.Atoi(last)
		id = int(id)
		fmt.Printf("Current notification URL: %s\n", *notification.Subject.URL)
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
	// TODO: replace with proper regexp, maybe have a list of keywords to look
	// for.
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
	logrus.Print("test")
}

func mergeCommit() error {
	return nil
}

// SetRepoSubTrue subscribes to all the repositories in a GitHub organization.
// TODO: implement black-listed repos.
func SetRepoSubTrue(ghDetails *GitHub) {
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

func addLabels() error {
	return nil
}
