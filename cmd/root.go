package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	urlProperty = "url"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "homewizard-prometheus-exporter",
		Short: "The homewizard prometheus metrics exporter",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
	}
	rootCmd.PersistentFlags().StringP(urlProperty, "u", "", "The URL of the Homewizard device (required)")

	rootCmd.AddCommand(createLocalUser)
	rootCmd.AddCommand(exportMetrics)
	return rootCmd
}

func initializeConfig(rootCmd *cobra.Command) error {
	viper.SetEnvPrefix("homewizard_prometheus_exporter")
	viper.AutomaticEnv()

	viper.BindPFlag(urlProperty, rootCmd.PersistentFlags().Lookup(urlProperty))
	return nil
}
