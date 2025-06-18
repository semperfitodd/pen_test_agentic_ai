package scanners

import (
	"os/exec"
)

type Nikto struct{}

func NewNikto() *Nikto { return &Nikto{} }

func (n *Nikto) Name() string { return "Nikto" }

func (n *Nikto) Run(target string) (string, error) {
	cmd := exec.Command("perl", "/opt/nikto/program/nikto.pl", "-host", target, "-port", "443", "-ssl")
	output, err := cmd.CombinedOutput()
	return string(output), err
}