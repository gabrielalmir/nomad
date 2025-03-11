package application

import (
	"fmt"
	"os"
	"os/exec"
)

type Pacman struct {
	Name string
}

func (p *Pacman) Update() error {
	cmd := exec.Command("pacman", "-Syy")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (p *Pacman) Install(packages []string) error {
	fmt.Println(packages)
	return nil
}

func NewPacman() *Pacman {
	return &Pacman{
		Name: "pacman",
	}
}
