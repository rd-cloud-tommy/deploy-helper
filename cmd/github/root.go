package github

import (
	"github.com/spf13/cobra"
)

var (
	inputToken        string
	inputOwner        string
	inputRepo         string
	inputTag          string
	inputProject      string
	inputDomain       string
	inputSlackWebhook string
	inputSlackChannel string
	dryRun            bool
)

// NewGithubCmd return newCommand
func NewGithubCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "github",
		Short: "github helper function",
	}

	command.PersistentFlags().StringVarP(&inputToken, "token", "t", "", "GitHub OAuth token")
	command.AddCommand(newReleaseNotifyCmd())

	return command
}
