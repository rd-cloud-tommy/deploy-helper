package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string
var githash string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Deploy Helper",
	Long:  `All software has versions. This is Deploy Helper's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deploy Helper \nversion: %s\ngithash: %s\n", version, githash)
	},
}
