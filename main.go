package main

import (
	"os"

	"github.com/gabrielalmir/nomad/infrastructure"
)

func main() {
	args := os.Args[1:]

	cli := infrastructure.CLI{
		Name:        "Nomad",
		Description: "Unique interface across different Linux distros",
	}

	if len(args) < 1 {
		cli.PrintUsage()
		return
	}

	cli.Execute(args)
}
