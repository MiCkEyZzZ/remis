package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func NewGitHubClient() (*github.Client, context.Context) {
	token := viper.GetString("token")
	if token == "" {
		fmt.Fprintln(os.Stderr, "GitHub token required: set GITHUB_TOKEN or --token")
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client, ctx
}
