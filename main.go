package main

import (
	"os"

	"github.com/gabrielalmir/nomad/ui"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		ui.PrintUsage()
	}
}
