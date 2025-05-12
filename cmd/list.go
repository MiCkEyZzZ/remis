package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var listCmd = &cobra.Command{
	Use:   "list [owner] [repo]",
	Short: "List issues for a repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo := args[0], args[1]

		token := viper.GetString("token")
		if token == "" {
			fmt.Fprintln(os.Stderr, "GitHub token required: set GITHUB_TOKEN or --token")
			os.Exit(1)
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		issues, _, err := client.Issues.ListByRepo(ctx, owner, repo, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching issues: %v\n", err)
			os.Exit(1)
		}

		for _, issue := range issues {
			fmt.Printf("#%d %s [%s]\n", issue.GetNumber(), issue.GetTitle(), issue.GetState())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
