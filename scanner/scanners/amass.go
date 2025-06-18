package scanners

import (
	"os/exec"
)

type Amass struct{}

func NewAmass() *Amass { return &Amass{} }

func (a *Amass) Name() string { return "Amass" }

func (a *Amass) Run(target string) (string, error) {
	cmd := exec.Command("amass", "enum", "-d", target)
	output, err := cmd.CombinedOutput()
	return string(output), err
}