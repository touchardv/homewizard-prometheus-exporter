package cmd

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/touchardv/homewizard-prometheus-exporter/pkg/homewizard"
)

var exportMetrics = &cobra.Command{
	Use: "export-metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := viper.GetString("url")
		token := viper.GetString("token")
		client := homewizard.NewAPIv2Client(url, homewizard.WithToken(token))
		reg := prometheus.NewRegistry()
		reg.Register(client)
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
		http.ListenAndServe(":8080", nil)
		return nil
	},
}

func init() {
	exportMetrics.Flags().StringP("token", "t", "", "The user authentication token")
	viper.BindPFlag("token", exportMetrics.Flags().Lookup("token"))
}
