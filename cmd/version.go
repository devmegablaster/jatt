package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var JattVersion string

func getJattVersion() string {
	if len(JattVersion) > 0 {
		return JattVersion
	}

	return "Unable to get Jatt version."
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Get the version of Jatt",
	Long:  `The version command provides information about the application's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getJattVersion())
	},
}
