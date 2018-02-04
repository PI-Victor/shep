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

func (g *GitHub) Validate(ctx context.Context) error {
	return nil
}

func NewGithubService(cfg *Config) *GitHub {
	return &GitHub{
		User:        cfg.GitHub.User,
		Token:       cfg.GitHub.Token,
		IgnoreRepos: cfg.GitHub.IgnoreRepos,
		Client:      cfg.GitHub.Client,
	}
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

type ghBotCommand func(context.Context, *GitHub, *prDetails) error

var cmdMapping = map[string]ghBotCommand{
	"merge": mergePR,
	"test":  runTestJob,
	"tag":   addLabels,
}

func newPRDetails(ctx context.Context, client *github.Client, notification *github.Notification) (*prDetails, error) {
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
	logrus.Debugf("last index: %d", lastIndex)
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

func getIssueDetails(ctx context.Context, client *github.Client, notification *github.Notification) error {
	logrus.Debug("trying to see if this is an issue...")
	logrus.Debugf("notification is :%#v", *notification.Reason)
	return nil
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

// Run watches the repositories in the organization that the bot is part
// of for events such as comments or PRs.
func (g *GitHub) Run(ctx context.Context) error {
	client := g.Client
	opt := &github.NotificationListOptions{All: true}
	//logrus.Debugf("This is the client %#v and these are the options: %#v and this is the context: %#v", g.Client, opt, ctx)
	notifications, _, err := client.Activity.ListNotifications(ctx, opt)
	if err != nil {
		return err
	}
	for _, notification := range notifications {
		prDetails, err := newPRDetails(ctx, g.Client, notification)
		if err != nil {
			logrus.Debugf("Failed to get new PR details: %#v", err)
			if err = getIssueDetails(ctx, g.Client, notification); err != nil {
				logrus.Debugf("We didn't get anything: %#v", err)
			}
			logrus.Debug("Continuing...")
			continue
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
				if err := checkComment(ctx, g, comment.GetBody(), prDetails); err != nil {
					logrus.Warningf("Failed to apply action %s", err)
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
		cmd, ok := cmdMapping[comm]
		if !ok {
			logrus.Debugf("Command %s not found", comm)
			return nil
		}
		if err := cmd(ctx, githubDetails, pr); err != nil {
			return err
		}
	}
	return nil
}

func commentPR(ctx context.Context, githubDetails *GitHub, pr *prDetails, msg *string) error {
	logrus.Debug("Testing PR...")
	commentBody := msg
	comm := github.IssueComment{
		Body: commentBody,
	}
	_, _, err := githubDetails.Client.Issues.CreateComment(ctx, pr.owner, pr.repo, pr.id, &comm)
	if err != nil {
		return err
	}
	return nil
}

func mergePR(ctx context.Context, githubDetails *GitHub, pr *prDetails) error {
	logrus.Debug("Merging PR...")
	msg := fmt.Sprintf("Merging PR branch %s into master", pr.branch)
	isMerged, _, err := githubDetails.Client.PullRequests.IsMerged(ctx, pr.owner, pr.repo, pr.id)
	if err != nil {
		return err
	}
	if !isMerged {
		// NOTE: this could be tricky if the response of the action is different
		// than successful. i don't expect an err everytime.
		_, _, err := githubDetails.Client.PullRequests.Merge(
			ctx, pr.owner, pr.repo, pr.id,
			"Merge based on [URL of CI/CD job]()",
			nil,
		)
		if err != nil {
			return err
		}
		commentPR(ctx, githubDetails, pr, &msg)
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

func addLabels(ctx context.Context, ghDetails *GitHub, pr *prDetails) error {
	return nil
}

func runTestJob(ctx context.Context, ghDetails *GitHub, pr *prDetails) error {
	msg := fmt.Sprintf("Running Test job")
	commentPR(ctx, ghDetails, pr, &msg)
	return nil
}
