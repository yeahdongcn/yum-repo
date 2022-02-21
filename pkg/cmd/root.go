package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yeahdongcn/yum-repo/pkg/repo"
)

var (
	repoPath string
	rootCmd  = &cobra.Command{
		Use:   "watcher",
		Short: "TODO",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			repo.Watch(repoPath)
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&repoPath, "repo", "r", "/repo", "TODO")
}

func Execute() error {
	return rootCmd.Execute()
}
