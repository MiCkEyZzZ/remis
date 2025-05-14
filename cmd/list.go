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
		perPage := viper.GetInt("per-page")
		if perPage <= 0 {
			perPage = 10
		}
		page := 1

		for {
			opts := &github.IssueListByRepoOptions{
				State: state,
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: perPage,
				},
			}

			items, _, err := client.Issues.ListByRepo(ctx, owner, repo, opts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching issues: %v\n", err)
				os.Exit(1)
			}

			if len(items) == 0 && page == 1 {
				fmt.Printf("No issues found for '%s/%s' with state='%s'\n", owner, repo, state)
				return
			} else if len(items) == 0 {
				fmt.Println("No more issues.")
				page--
				continue
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
			fmt.Fprintln(w, "Number\tState\tTitle")
			for _, issue := range items {
				st := issue.GetState()
				var stCol string
				if st == "open" {
					stCol = color.New(color.FgGreen).Sprint("● OPEN")
				} else {
					stCol = color.New(color.FgRed).Sprint("● CLOSED")
				}

				fmt.Fprintf(w, "#%d\t%s\t%s\n",
					issue.GetNumber(),
					stCol,
					issue.GetTitle(),
				)
			}
			w.Flush()

			fmt.Print("\n(n = next page, p = previous page, q = quit): ")
			var input string
			fmt.Scanln(&input)

			switch input {
			case "n":
				page++
			case "p":
				if page > 1 {
					page--
				} else {
					fmt.Println("Already at the first page.")
				}
			case "q", "":
				return
			default:
				fmt.Println("Invalid input. Use n (next), p (previous), or q (quit).")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("state", "s", "open", "Filter by state: open, closed, all")
	listCmd.Flags().Int("per-page", 10, "Number of issues per page")
	viper.BindPFlag("state", listCmd.Flags().Lookup("state"))
	viper.BindPFlag("per-page", listCmd.Flags().Lookup("per-page"))
}
