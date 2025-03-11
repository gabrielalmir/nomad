package application

import (
	"errors"
	"fmt"
	"log"
)

type Packager struct {
	Name   string
	IsRoot bool
}

func (p *Packager) Update() {
	fmt.Println("Updating system")

	var err error

	switch p.Name {
	case "pacman":
		pacman := NewPacman()
		err = pacman.Update()
	default:
		err = errors.New("apt isn't yet implemented")
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func NewPackager(name string, isRoot bool) *Packager {
	return &Packager{
		Name:   name,
		IsRoot: isRoot,
	}
}
