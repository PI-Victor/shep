package service

import (
	"context"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/PI-Victor/shep/pkg/fs"
)

// NewGitHubClient returns a valid instance of a new GitHub client.
func NewGitHubClient(cfg *fs.Config) {
	ctx := context.Background()
	tokenSrc := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: cfg.GitHubToken,
		},
	)
	tc := oauth2.NewClient(ctx, tokenSrc)
	client := github.NewClient(tc)
	repos, _, err := client.Repositories.List(ctx, cfg.GitHubUser, nil)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(repos)
}
