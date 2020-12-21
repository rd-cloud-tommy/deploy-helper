package cmd

import (
	"deploy-helper/services/aws"

	"github.com/spf13/cobra"
)

var allowTraffic = &cobra.Command{
	Use:   "allow-traffic",
	Short: "register instance to target group",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := aws.New()
		if err != nil {
			panic(err)
		}

		err = c.RegisterInstanceSelf()
		if err != nil {
			panic(err)
		}
	},
}
