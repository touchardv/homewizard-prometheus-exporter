package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/touchardv/homewizard-prometheus-exporter/pkg/homewizard"
)

var createLocalUser = &cobra.Command{
	Use: "create-user",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := viper.GetString("url")
		username, _ := cmd.Flags().GetString("username")
		client := homewizard.NewAPIv2Client(url)
		return client.CreateLocalUser(username)
	},
}

func init() {
	createLocalUser.Flags().StringP("username", "", "", "The name of the user to register in the Homewizard device (required)")
	createLocalUser.MarkFlagRequired("username")
}
