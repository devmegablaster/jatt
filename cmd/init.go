package cmd

import (
	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/writer"
	"github.com/spf13/cobra"
)

var initJatt = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfigTemplate()

		var projectName string
		if len(args) > 0 {
			projectName = args[0]
		} else {
			projectName = "jatt"
		}

		w := writer.NewWriter(cfg)
		w.InitJatt(projectName)
	},
}
