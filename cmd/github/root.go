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
	inputSlackWebhook string
	inputSlackChannel string
)

// NewGithubCmd returns a new GitHub command
func NewGithubCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "github",
		Short: "GitHub helper functions",
	}

	command.PersistentFlags().StringVarP(&inputToken, "token", "t", "", "GitHub OAuth token")
	command.AddCommand(newReleaseNotifyCmd())

	return command
}
