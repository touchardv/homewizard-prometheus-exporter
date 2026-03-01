package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	urlProperty   = "url"
	tokenProperty = "token"
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
	rootCmd.MarkPersistentFlagRequired(urlProperty)
	viper.BindPFlag(urlProperty, rootCmd.PersistentFlags().Lookup(urlProperty))

	rootCmd.AddCommand(createLocalUser)
	rootCmd.AddCommand(exportMetrics)
	rootCmd.AddCommand(listUsers)
	return rootCmd
}

func initializeConfig(cmd *cobra.Command) error {
	viper.SetEnvPrefix("homewizard_prometheus_exporter")
	viper.AutomaticEnv()

	bindFlags(cmd, viper.GetViper())

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// If the user hasn't set the flag, use the value from Viper (config/env)
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
