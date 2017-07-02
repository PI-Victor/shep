package services

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// TODO: this code needs to use DI. since it's not testable at all.
// isolating component deps would make this easy to test.

var (
	// lastCheck keeps the time that we checked for new notifications. if the
	// comment of the PR is older than lastCheck, we skip it.
	lastCheck time.Time
)

// GitHub holds specific information that is used for GitHub integration.
type GitHub struct {
	User  *github.User `json:"-"`
	Token string       `json:"token"`

	// A list of URLs that the bot can ignore.
	IgnoreRepos []string `json:"ignoreRepos,omitempty"`
	// Holds the client instance details. Internal only.
	Client *github.Client `json:"-"`
}

type prDetails struct {
	state    string
	branch   string
	isMaster bool
	merged   bool
	owner    string
	repo     string
	id       int
}

func newPRDetails(
	ctx context.Context,
	client *github.Client,
	notification *github.Notification,
) (*prDetails, error) {
	// TODO: replace this with a proper group regexp.
	var (
		urlParts       = strings.Split(*notification.Subject.URL, "/")
		repository     = notification.Repository.GetName()
		repoOwner      = urlParts[4]
		isMasterBranch bool
	)
	// It's safe to assume that, since the API is standard these indexes will
	// always extract the correct data.
	lastIndex, err := strconv.Atoi(urlParts[7])
	if err != nil {
		return nil, err
	}
	prID := int(lastIndex)
	rawPR, _, err := client.PullRequests.Get(ctx, repoOwner, repository, prID)
	if err != nil {
		return nil, err
	}
	if *rawPR.Head.Ref == "master" {
		isMasterBranch = true
	}
	return &prDetails{
		owner:    repoOwner,
		repo:     repository,
		isMaster: isMasterBranch,
		branch:   *rawPR.Head.Ref,
		state:    *rawPR.State,
		merged:   *rawPR.Merged,
		id:       prID,
	}, nil
}

// NewGitHubClient returns a valid instance of a new GitHub client.
func NewGitHubClient(ctx context.Context, cfg *Config) error {
	tokenSrc := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: cfg.GitHub.Token,
		},
	)
	// TODO: implement HTTPS based authentication.
	tc := oauth2.NewClient(ctx, tokenSrc)
	client := github.NewClient(tc)
	cfg.GitHub.Client = client
	authClient, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return err
	}

	cfg.GitHub.User = authClient
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
		prDetails, err := newPRDetails(ctx, ghDetails.Client, notification)
		if err != nil {
			return err
		}
		comments, _, err := client.Issues.ListComments(
			ctx,
			prDetails.owner,
			prDetails.repo,
			prDetails.id,
			nil,
		)
		if err != nil {
			logrus.Warning(err)
			continue
		}
		for _, comment := range comments {
			if lastCheck.Before(comment.GetCreatedAt()) {
				if err := checkComment(ctx, ghDetails, comment.GetBody(), prDetails); err != nil {
					logrus.Warningf("failed to apply action %s", err)
				}
			}
		}
	}
	// NOTE: this will not be needed anymore after MarkRepositoryNotificationsRead
	lastCheck = time.Now()
	return nil
}

func checkComment(ctx context.Context, githubDetails *GitHub, body string, pr *prDetails) error {
	// TODO: replace with proper regexp, maybe have a list of keywords to look
	// for.
	match, err := regexp.Match("[[a-zA-Z]+]", []byte(body))
	if err != nil {
		return err
	}
	if match {
		comm := strings.Trim(body, "[]")
		if comm == "test" {
			if err := commentPR(ctx, githubDetails, pr); err != nil {
				return err
			}
		}
		if comm == "merge" {
			if err := mergePR(ctx, githubDetails, pr); err != nil {
				return err
			}
		}
	}
	return nil
}

func commentPR(ctx context.Context, githubDetails *GitHub, pr *prDetails) error {
	logrus.Debug("Testing PR...")
	commentBody := "Testing URL [To be filled in - CI/CD URL]"
	comm := github.IssueComment{
		Body: &commentBody,
	}
	_, _, err := githubDetails.Client.Issues.CreateComment(ctx, pr.owner, pr.repo, pr.id, &comm)
	if err != nil {
		return err
	}
	return nil
}

func mergePR(ctx context.Context, githubDetails *GitHub, pr *prDetails) error {
	logrus.Debug("Merging PR...")
	isMerged, _, err := githubDetails.Client.PullRequests.IsMerged(ctx, pr.owner, pr.repo, pr.id)
	if err != nil {
		return err
	}
	if !isMerged {
		// NOTE: this could be tricky if the response of the action is different
		// than successful. i don't expect an err everytime.
		_, _, err := githubDetails.Client.PullRequests.Merge(
			ctx, pr.owner, pr.repo, pr.id,
			"Merge based on [URL of PR to be filled in]",
			nil,
		)
		if err != nil {
			return err
		}
		pr.merged = true
	}
	if err := deleteBranch(ctx, githubDetails, pr); err != nil {
		logrus.Debugf("Failed to delete branch: %+v", err)
	}
	return nil
}

// NOTE: is it a bad idea to delete branches that were used to merge a PR?
func deleteBranch(ctx context.Context, githubDetails *GitHub, pr *prDetails) error {
	logrus.Debugf("the pr is merged: %#v", pr)
	if !pr.isMaster && pr.merged {
		logrus.Debugf("The pr ref in deleteBranch is :%s", pr.branch)
		_, err := githubDetails.Client.Git.DeleteRef(ctx, pr.owner, pr.repo,
			strings.Replace("heads/"+pr.branch, "#", "%23", -1),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func closePullRequest(ctx context.Context, githubDetails *GitHub, pr *prDetails) error {
	logrus.Debug("Closing PR...")
	_, err := githubDetails.Client.Activity.MarkRepositoryNotificationsRead(
		ctx,
		pr.owner,
		pr.repo,
		time.Now(),
	)
	if err != nil {
		return err
	}
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
			sub, _, err := client.Activity.SetRepositorySubscription(
				ctx,
				org.GetLogin(),
				repo.GetName(),
				&subscription,
			)
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
