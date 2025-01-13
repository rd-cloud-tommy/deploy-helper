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
		Run:   runReleaseNotifyCmd,
	}

	command.Flags().StringVarP(&inputOwner, "owner", "", "", "github owner(required)")
	command.Flags().StringVarP(&inputRepo, "repo", "", "", "github repo(required)")
	command.Flags().StringVarP(&inputTag, "tag", "", "", "github tag(required)")
	command.Flags().StringVarP(&inputProject, "project", "", "", "project")
	command.Flags().StringVarP(&inputDomain, "domain", "", "", "service domain")
	command.Flags().StringVarP(&inputSlackWebhook, "slackWebhook", "", "", "slack webhook")
	command.Flags().StringVarP(&inputSlackChannel, "slackChannel", "", "", "slack channel")
	command.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry-run mode for local testing")
	command.MarkFlagRequired("owner")
	command.MarkFlagRequired("repo")
	command.MarkFlagRequired("tag")

	return command
}

func runReleaseNotifyCmd(cmd *cobra.Command, args []string) {
	githubClient := github.New(inputToken)
	repoRelease, err := githubClient.GetReleaseByTag(inputOwner, inputRepo, inputTag)

	if err != nil {
		panic(err)
	}

	if inputSlackChannel != "" && inputSlackWebhook != "" {

		// Set inputProject to inputRepo if inputProject is not set
		if inputProject == "" {
			inputProject = inputRepo
		}
	
		text := fmt.Sprintf(
			"%s deployment successful. \nProject: %s \nVersion: %s \nDomain: %s \n<%s|Click here> for details! \n\n%s",
			inputRepo,
			inputProject,
			inputTag,
			inputDomain,
			repoRelease.GetHTMLURL(),
			repoRelease.GetBody(),
		)

		if dryRun {
			fmt.Println(text)
		} else {
			sendSlackNotification(text)
		}
	} else {
		fmt.Println(repoRelease.GetBody())
	}
}

func sendSlackNotification(text string) {
	msg := &slack.WebhookMessage{
		Channel: inputSlackChannel,
		Text:    text,
	}

	err := slack.PostWebhook(inputSlackWebhook, msg)

	if err != nil {
		fmt.Printf("Failed to send slack notification: %v\n", err)
	}
}
