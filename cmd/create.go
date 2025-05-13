package cmd

import (
	"fmt"
	"os"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/cobra"
)

var createBodyFile string

var createCmd = &cobra.Command{
	Use:   "create [owner] [repo] [title]",
	Short: "Create a new issue",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo, title := args[0], args[1], args[2]

		body := ""
		if createBodyFile != "" {
			data, err := os.ReadFile(createBodyFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading body file: %v\n", err)
				os.Exit(1)
			}
			body = string(data)
		}

		client, ctx := NewGitHubClient()

		issueRequest := &github.IssueRequest{Title: &title, Body: &body}
		issue, _, err := client.Issues.Create(ctx, owner, repo, issueRequest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating issue: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created issue #%d %s\n", *issue.Number, *issue.HTMLURL)
	},
}

func init() {
	createCmd.Flags().StringVarP(&createBodyFile, "body-file", "b", "", "path to file for issue body")
	rootCmd.AddCommand(createCmd)
}
