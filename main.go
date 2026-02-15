package main

import (
	"os"

	"github.com/touchardv/homewizard-prometheus-exporter/cmd"
)

func main() {
	cmd := cmd.NewRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
