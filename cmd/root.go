package cmd

import (
	"os"

	"github.com/docker/docker/client"
	"github.com/localghost/swarmer/cmd/container"
	"github.com/localghost/swarmer/cmd/node"
	"github.com/spf13/cobra"
)

type rootFlags struct {
	host string
}

func NewRootCommand() *cobra.Command {
	flags := rootFlags{}

	cmd := &cobra.Command{
		Use: "swarmer [OPTIONS] COMMAND",
		PreRun: func(cmd *cobra.Command, args []string) {
			if flags.host == "" {
				flags.host = os.Getenv("DOCKER_HOST")
				if flags.host == "" {
					flags.host = client.DefaultDockerHost
				}
			}
		},
	}

	cmd.Flags().StringVarP(&flags.host, "host", "H", "", "Manager host.")

	addCommands(cmd)

	return cmd
}

func addCommands(root *cobra.Command) {
	root.AddCommand(container.NewPsCommand())
	root.AddCommand(node.NewNodeCommand())
}
