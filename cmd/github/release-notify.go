package github

import (
	"deploy-helper/components/aws/github"
	"fmt"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
)

func newReleaseNotifyCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "release-notify",
		Short: "get release body and set notify",
		Run: func(cmd *cobra.Command, args []string) {
			githubClient := github.New(inputToken)
			repoRelease, err := githubClient.GetReleaseByTag(inputOwner, inputRepo, inputTag)
			if err != nil {
				panic(err)
			}

			// sent message to slack
			if inputSlackChannel != "" && inputSlackWebhook != "" {
				msg := &slack.WebhookMessage{
					Channel: inputSlackChannel,
					Text:    fmt.Sprintf("%s deploy success. version: %s <%s|Click here> for details!\n%s", inputRepo, inputTag, repoRelease.GetHTMLURL(), repoRelease.GetBody()),
				}
				slack.PostWebhook(inputSlackWebhook, msg)
			} else {
				fmt.Println(repoRelease.GetBody())
			}
		},
	}

	command.Flags().StringVarP(&inputOwner, "owner", "", "", "github owner(required)")
	command.Flags().StringVarP(&inputRepo, "repo", "", "", "github repo(required)")
	command.Flags().StringVarP(&inputTag, "tag", "", "", "github tag(required)")
	command.Flags().StringVarP(&inputSlackWebhook, "slackWebhook", "", "", "slack webhook")
	command.Flags().StringVarP(&inputSlackChannel, "slackChannel", "", "", "slack channel")
	command.MarkFlagRequired("owner")
	command.MarkFlagRequired("repo")
	command.MarkFlagRequired("tag")

	return command
}
