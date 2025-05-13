package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "remis",
	Short: "Remis - CLI for managing GitHub Issues",
	Long:  "A terminal tool to create, list, and mange GitHub issues using Go, Cobra and Viper.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.BindEnv("token", "GITHUB_TOKEN")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_ = viper.SafeWriteConfig()
		} else {
			fmt.Fprintln(os.Stderr, "Error reading config file:", err)
			os.Exit(1)
		}
	}

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is ./config.yaml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().StringP("token", "t", "", "GitHub OAuth token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}
