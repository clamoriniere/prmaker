package prmaker

import (
	"context"
	"fmt"

	"github.com/clamoriniere/prmaker/cmd/prmaker/options"
	"github.com/clamoriniere/prmaker/pkg/github"
	"github.com/spf13/cobra"
)

type createOptions struct {
	Owner    string
	RepoName string

	global *options.Global
}

func NewCreateCmd(global *options.Global) *cobra.Command {
	opts := createOptions{
		global: global,
	}
	createCmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{},
		Short:   "create a pull request",
		//Args:    cobra.ExactArgs(1),
		RunE: opts.run,
	}

	createCmd.Flags().StringVarP(&opts.Owner, "owner", "", "", "github repository owner")
	createCmd.Flags().StringVarP(&opts.RepoName, "repository", "", "", "github repository name")

	return createCmd
}

func (opts *createOptions) run(cmd *cobra.Command, args []string) error {
	fmt.Fprintf(cmd.OutOrStdout(), "TOKEN: %s, OWNER: %s, REPO: %s\n", opts.global.GithubToken, opts.Owner, opts.RepoName)

	client := github.NewClient(opts.global.GithubToken)
	desc, err := client.FetchRepoDescription(context.TODO(), opts.Owner, opts.RepoName)
	if err != nil {
		return err
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Description: %s\n", desc)

	file, err := client.FetchFile(context.TODO(), opts.Owner, opts.RepoName, "HEAD", "README.md")
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "File: %s\n", file.Blob.Text)

	return nil
}
