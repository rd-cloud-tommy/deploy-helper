package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func newReleaseMsgCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "release-msg",
		Short: "get release note",
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
			fmt.Println(repoRelease.GetBody())
		},
	}

	command.Flags().StringVarP(&inputOwner, "owner", "", "", "github owner(required)")
	command.Flags().StringVarP(&inputRepo, "repo", "", "", "github repo(required)")
	command.Flags().StringVarP(&inputTag, "tag", "", "", "github tag(required)")
	command.MarkFlagRequired("owner")
	command.MarkFlagRequired("repo")
	command.MarkFlagRequired("tag")

	return command
}
