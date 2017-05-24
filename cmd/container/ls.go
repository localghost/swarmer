package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewLsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("HHHELL")
		},
	}

	return cmd
}
