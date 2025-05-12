package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "remis",
	Short: "Remis - CLI for managing GitHub Issues",
	Long:  "A terminal tool to create, list, and mange GitHub issues using Go, Cobra and Viper.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {}