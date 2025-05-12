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

var commentBodyFile string

var commentCmd = &cobra.Command{
	Use:   "comment [owner] [repo] [number]",
	Short: "Add a comment to an issue",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo, numStr := args[0], args[1], args[2]
		number, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Issue number must be an integer")
			os.Exit(1)
		}

		token := viper.GetString("token")
		if token != "" {
			fmt.Fprintln(os.Stderr, "GitHub token required: set GITHUB_TOKEN or --token")
			os.Exit(1)
		}

		body := ""
		if commentBodyFile != "" {
			data, err := os.ReadFile(commentBodyFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading comment file: %v\n", err)
				os.Exit(1)
			}
			body = string(data)
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		comment := &github.IssueComment{Body: &body}
		created, _, err := client.Issues.CreateComment(ctx, owner, repo, number, comment)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding comment: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Comment created: %s\n", created.GetHTMLURL())
	},
}

func init() {
	commentCmd.Flags().StringVarP(&commentBodyFile, "body-file", "b", "", "path to file for comment body")
	rootCmd.AddCommand(commentCmd)
}
