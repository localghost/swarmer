package container

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func NewPsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ps",
		Run: func(cmd *cobra.Command, args []string) {
			listTasks()
		},
	}

	return cmd
}

func listTasks() {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	tasks, err := cli.TaskList(context.Background(), types.TaskListOptions{})
	if err != nil {
		log.Fatal(nil)
	}
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		log.Fatal(nil)
	}
	nodeInfos := make(map[string]*swarm.Node)
	for _, node := range nodes {
		nodeInfos[node.ID] = &node
	}
	serviceInfos := serviceList(cli)
	for _, task := range tasks {
		fmt.Printf(
			"Task: %s.%s, container: %s, node: %s\n",
			serviceInfos[task.ServiceID].Spec.Name,
			task.ID,
			task.Status.ContainerStatus.ContainerID,
			nodeInfos[task.NodeID].Status.Addr,
		)
	}
}

func serviceList(cli *client.Client) (serviceInfos map[string]*swarm.Service) {
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	serviceInfos = make(map[string]*swarm.Service)
	for _, service := range services {
		serviceInfos[service.ID] = &service
	}
	return
}
