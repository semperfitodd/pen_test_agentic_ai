package scanners

import (
	"os/exec"
)

type Nmap struct{}

func NewNmap() *Nmap { return &Nmap{} }

func (n *Nmap) Name() string { return "Nmap" }

func (n *Nmap) Run(target string) (string, error) {
	cmd := exec.Command("nmap", "-p", "80,443", "-sV", "-T4", "-Pn", target)
	output, err := cmd.CombinedOutput()
	return string(output), err
}