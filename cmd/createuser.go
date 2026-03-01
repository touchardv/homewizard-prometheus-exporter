package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/touchardv/homewizard-prometheus-exporter/pkg/homewizard"
)

const usernameProperty = "username"

var createLocalUser = &cobra.Command{
	Use: "create-user",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := viper.GetString(urlProperty)
		username, _ := cmd.Flags().GetString(usernameProperty)
		client := homewizard.NewAPIv2Client(url)
		return client.CreateLocalUser(username)
	},
}

func init() {
	createLocalUser.Flags().StringP(usernameProperty, "", "", "The name of the user to register in the Homewizard device (required)")
	createLocalUser.MarkFlagRequired(usernameProperty)
}
