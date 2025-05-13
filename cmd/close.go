package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close [owner] [repo] [number]",
	Short: "Close an existing issue",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo, numStr := args[0], args[1], args[2]
		number, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Issue number must be an integer")
			os.Exit(1)
		}

		client, ctx := NewGitHubClient()

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
