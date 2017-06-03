package container

import (
	"github.com/spf13/cobra"
)

func NewContainerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "container [OPTIONS] COMMAND",
	}

	cmd.AddCommand(NewPsCommand())

	return cmd
}
