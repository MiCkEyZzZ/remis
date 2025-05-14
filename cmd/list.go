package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
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

		items, _, err := client.Issues.ListByRepo(ctx, owner, repo, opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching issues: %v\n", err)
			os.Exit(1)
		}

		if len(items) == 0 {
			msg := fmt.Sprintf("No issues found for '%s/%s' with state='%s'", owner, repo, state)
			fmt.Println(msg)
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
		fmt.Fprintln(w, "Number\tState\tTitle")
		for _, issue := range items {
			// Цветное состояние
			st := issue.GetState()
			var stCol string
			if st == "open" {
				stCol = color.New(color.FgGreen).Sprint("● OPEN")
			} else {
				stCol = color.New(color.FgRed).Sprint("● CLOSED")
			}

			// Выводим: метки в квадратных скобках
			fmt.Fprintf(w, "#%d\t%s\t%s\n",
				issue.GetNumber(),
				stCol,
				issue.GetTitle(),
			)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("state", "s", "open", "Filter by state: open, closed, all")
	viper.BindPFlag("state", listCmd.Flags().Lookup("state"))
}
