package lib

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"context"
	"github.com/urfave/cli"
	"os"
	"net/http"
)

func GetGithubClient(c *cli.Context) *github.Client {
	var tc *http.Client
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		)
		tc = oauth2.NewClient(context.Background(), ts)
	}
	return github.NewClient(tc)
}

