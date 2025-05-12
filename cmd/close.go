package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var closeCmd = &cobra.Command{
	Use:   "close [owner] [repo] [number]",
	Short: "Close an existing issue",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo, numStr := args[0], args[1], args[2]
		number, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Issue number must be ab integer")
			os.Exit(1)
		}

		token := viper.GetString("token")
		if token == "" {
			fmt.Fprintln(os.Stderr, "GitHub token required: set GITHUB_TOKEN or --token")
			os.Exit(1)
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		state := "closed"
		issueRequest := &github.IssueRequest{State: &state}
		issue, _, err := client.Issues.Edit(ctx, owner, repo, number, issueRequest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error closing issue: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Closed issue #%d %s\n", *issue.Number, issue.GetHTMLURL())
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
