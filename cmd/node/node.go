package node

import (
	"github.com/spf13/cobra"
)

func NewNodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "node [OPTIONS] COMMAND",
	}

	cmd.AddCommand(NewLsCommand())

	return cmd
}
