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

var createBodyFile string

var createCmd = &cobra.Command{
	Use:   "create [owner] [repo] [title]",
	Short: "Create a new issue",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo, title := args[0], args[1], args[2]

		token := viper.GetString("token")
		if token == "" {
			fmt.Fprintln(os.Stderr, "GitHub token required: set GITHUB_TOKEN or --token")
			os.Exit(1)
		}

		body := ""
		if createBodyFile != "" {
			data, err := os.ReadFile(createBodyFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading body file: %v\n", err)
				os.Exit(1)
			}
			body = string(data)
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

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
