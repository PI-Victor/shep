package services

import (
	"context"
	_ "fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var (
	// lastCheck keeps the time that we checked for new notifications. if the
	// comment of the PR is older than lastCheck, we skip it.
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

type prDetails struct {
	org  string
	repo string
	id   int
}

func newPRDetails(notification *github.Notification) (*prDetails, error) {
	// TODO: replace this with a proper group regexp.
	urlParts := strings.Split(*notification.Subject.URL, "/")
	repository := notification.Repository.GetName()
	// It's safe to assume that, since the API is standard these indexes will
	// always extract the correct data.
	organisation := urlParts[4]
	lastIndex, err := strconv.Atoi(urlParts[7])
	if err != nil {
		return nil, err
	}

	return &prDetails{
		org:  organisation,
		repo: repository,
		id:   int(lastIndex),
	}, nil
}

// NewGitHubClient returns a valid instance of a new GitHub client.
func NewGitHubClient(ctx context.Context, cfg *Config) error {
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
func WatchRepos(ctx context.Context, ghDetails *GitHub) error {
	client := ghDetails.Client
	opt := &github.NotificationListOptions{All: true}

	notifications, _, err := client.Activity.ListNotifications(ctx, opt)
	if err != nil {
		return err
	}
	for _, notification := range notifications {
		prDetails, err := newPRDetails(notification)
		if err != nil {
			return err
		}
		comments, _, err := client.Issues.ListComments(ctx, prDetails.org, prDetails.repo, prDetails.id, nil)
		if err != nil {
			logrus.Warning(err)
			continue
		}
		for _, comment := range comments {
			if lastCheck.Before(comment.GetCreatedAt()) {
				if err := checkComment(ctx, comment.GetBody(), prDetails); err != nil {
					logrus.Warningf("failed to apply action %s", err)
				}
			}
		}
	}
	// TODO: this is not really needed or used.
	lastCheck = time.Now()
	return nil
}

func checkComment(ctx context.Context, body string, pr *prDetails) error {
	// TODO: replace with proper regexp, maybe have a list of keywords to look
	// for.
	match, err := regexp.Match("[[a-zA-Z]+]", []byte(body))
	if err != nil {
		return err
	}
	if match {
		comm := strings.Trim(body, "[]")
		if comm == "test" {
			if err := commentPR(ctx, pr); err != nil {
				return err
			}
		}
		if comm == "merge" {
			if err := mergeCommit(ctx, pr); err != nil {
				return err
			}
		}
	}
	return nil
}

func commentPR(ctx context.Context, pr *prDetails) error {

	return nil
}

func mergeCommit(ctx context.Context, pr *prDetails) error {
	logrus.Print("Merging PR...")
	return nil
}

// SetRepoSubTrue subscribes to all the repositories in a GitHub organization.
// TODO: implement black-listed repos.
func SetRepoSubTrue(ctx context.Context, ghDetails *GitHub) {
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

func addLabels(ctx context.Context, ghDetails *GitHub) error {
	return nil
}
