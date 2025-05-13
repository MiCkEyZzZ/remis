package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [owner] [repo]",
	Short: "List issues for a repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo := args[0], args[1]

		client, ctx := NewGitHubClient()

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
