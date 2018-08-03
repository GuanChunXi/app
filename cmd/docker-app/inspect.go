package main

import (
	"github.com/docker/app/lib/dockerapp"
	"github.com/docker/app/internal/render"
	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
func inspectCmd(dockerCli command.Cli) *cobra.Command {
	return &cobra.Command{
		Use:   "inspect [<app-name>]",
		Short: "Shows metadata and settings for a given application",
		Args:  cli.RequiresMaxArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := dockerapp.Load(firstOrEmpty(args))
			if err != nil {
				return err
			}
			return render.Inspect(dockerCli.Out(), app)
		},
	}
}
