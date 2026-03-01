package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/touchardv/homewizard-prometheus-exporter/pkg/homewizard"
)

var listUsers = &cobra.Command{
	Use:   "list-users",
	Short: "List registered user(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := viper.GetString(urlProperty)
		token := viper.GetString(tokenProperty)
		client := homewizard.NewAPIv2Client(url, homewizard.WithToken(token))
		users, err := client.ListUsers()
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "% 20s %s\n", "NAME", "CURRENT")
		for _, user := range users {
			fmt.Fprintf(os.Stdout, "% 20s %t\n", user.Name, user.Current)
		}
		return nil
	},
}

func init() {
	listUsers.Flags().StringP(tokenProperty, "t", "", "The user authentication token")
	listUsers.MarkFlagRequired(tokenProperty)
	viper.BindPFlag(tokenProperty, listUsers.Flags().Lookup(tokenProperty))
}
