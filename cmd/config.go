package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure remis settings",
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key, value := args[0], args[1]

		if key == "token" {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: value})
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)

			user, _, err := client.Users.Get(ctx, "")
			if err != nil {
				fmt.Println("❌ Invalid token or network error:", err)
				return
			}
			fmt.Printf("✅ Token is valid. Authenticated as: %s\n", *user.Login)
		}

		viper.Set(key, value)

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Error writing config:", err)
			return
		}
		fmt.Printf("✔️ Set %s = %s\n", key, value)
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	rootCmd.AddCommand(configCmd)
}
