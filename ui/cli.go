package ui

import "fmt"

func PrintUsage() {
	fmt.Println("Usage: nomad <operation>")
	fmt.Println("Operations:")
	fmt.Println("\tnomad install <package(s)>")
	fmt.Println("\tnomad remove <package(s)>")
	fmt.Println("\tnomad update")
	fmt.Println("\tnomad upgrade")
	fmt.Println("\tnomad {-h --help}")
}
