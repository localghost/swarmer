package main

import (
	"os"

	"github.com/localghost/swarmer/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
