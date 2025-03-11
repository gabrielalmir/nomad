package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/gabrielalmir/nomad/application"
)

type CLI struct {
	Name        string
	Description string
}

func isRoot() bool {
	return os.Geteuid() == 0
}

func (c *CLI) Execute(args []string) {
	isUserRoot := isRoot()
	packager := application.NewPackager("pacman", isUserRoot)
	command := args[0]

	switch command {
	case "update":
		packager.Update()
	default:
		log.Fatalln("Not yet implemented")
	}
}

func (c *CLI) PrintUsage() {
	fmt.Println(c.Name, "-", c.Description)
	fmt.Println("Usage: nomad <operation>")
	fmt.Println("Operations:")
	fmt.Println("\tnomad install <package(s)>")
	fmt.Println("\tnomad remove <package(s)>")
	fmt.Println("\tnomad update")
	fmt.Println("\tnomad upgrade")
	fmt.Println("\tnomad {-h --help}")
}
