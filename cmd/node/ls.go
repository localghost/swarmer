package node

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func NewLsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			listNodes()
		},
	}

	return cmd
}

func listNodes() {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		log.Fatal(nil)
	}
	for _, node := range nodes {
		if node.ManagerStatus != nil {
			host, _, err := net.SplitHostPort(node.ManagerStatus.Addr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Manager: %s\n", host)
		} else {
			fmt.Printf("Worker: %s\n", node.Status.Addr)
		}
	}
}
