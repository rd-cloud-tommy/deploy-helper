package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func newReleaseNotifyCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "release-notify",
		Short: "get release body and set notify",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: inputToken},
			)
			tc := oauth2.NewClient(ctx, ts)

			client := github.NewClient(tc)

			repoRelease, _, err := client.Repositories.GetReleaseByTag(ctx, inputOwner, inputRepo, inputTag)
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
