package cmd

import (
	"fmt"
	"os"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list [owner] [repo]",
	Short: "List issues for a repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo := args[0], args[1]

		client, ctx := NewGitHubClient()
		state := viper.GetString("state")
		opts := &github.IssueListByRepoOptions{
			State: state,
		}
		issues, _, err := client.Issues.ListByRepo(ctx, owner, repo, opts)

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

	listCmd.Flags().StringP("state", "s", "open", "Filter by state: open, closed, all")
	viper.BindPFlag("state", listCmd.Flags().Lookup("state"))
}
