package cmd

import (
	"deploy-helper/services/aws"

	"github.com/spf13/cobra"
)

var blockTraffic = &cobra.Command{
	Use:   "block-traffic",
	Short: "deregister instance of target group",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := aws.New()
		if err != nil {
			panic(err)
		}

		err = c.DeregisterInstanceSelf()
		if err != nil {
			panic(err)
		}
	},
}
