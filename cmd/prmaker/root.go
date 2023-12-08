package prmaker

import (
	"fmt"
	"os"

	"github.com/clamoriniere/prmaker/cmd/prmaker/options"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	opts := &options.Global{}
	rootCmd := &cobra.Command{
		Use:   "prmaker",
		Short: "prmaker - a simple CLI to create a pull request from local files",
		Long:  `prmaker ease the creation of basic pull request without checkout the repository`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	rootCmd.PersistentFlags().StringVarP(&opts.GithubToken, "token", "", "", "github API token")

	rootCmd.AddCommand(NewCreateCmd(opts))

	return rootCmd
}

func Execute() {
	root := NewRootCmd()
	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
