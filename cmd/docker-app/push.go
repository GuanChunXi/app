package main

import (
	"github.com/docker/app/lib/dockerapp"
	"github.com/docker/cli/cli"
	"github.com/spf13/cobra"
)

type pushOptions struct {
	namespace string
	tag       string
}

func pushCmd() *cobra.Command {
	var opts pushOptions
	cmd := &cobra.Command{
		Use:   "push [<app-name>]",
		Short: "Push the application to a registry",
		Args:  cli.RequiresMaxArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := dockerapp.Load(firstOrEmpty(args))
			if err != nil {
				return err
			}
			return dockerapp.ToImageWithDefaults(app, opts.namespace, opts.tag)
		},
	}
	cmd.Flags().StringVar(&opts.namespace, "namespace", "", "namespace to use (default: namespace in metadata)")
	cmd.Flags().StringVarP(&opts.tag, "tag", "t", "", "tag to use (default: version in metadata)")
	return cmd
}
