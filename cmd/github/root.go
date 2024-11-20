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

// NewGithubCmd return newCommand
func NewGithubCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "github",
		Short: "github helper function",
	}

	command.PersistentFlags().StringVarP(&inputToken, "token", "t", "", "github oauth token")

	command.AddCommand(newReleaseNotifyCmd())
	return command
}
