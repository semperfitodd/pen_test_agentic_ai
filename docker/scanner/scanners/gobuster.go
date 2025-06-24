package scanners

import (
	"fmt"
	"os/exec"
)

type Gobuster struct{}

func NewGobuster() *Gobuster { return &Gobuster{} }

func (g *Gobuster) Name() string { return "Gobuster" }

func (g *Gobuster) Run(target string) (string, error) {
	cmd := exec.Command(
		"gobuster", "dir",
		"-u", "https://"+target,
		"-w", "/usr/share/wordlists/dirb/common.txt",
		"--timeout", "10s",
		"--threads", "10",
		"--no-error",
	)

	output, err := cmd.CombinedOutput()

	// Always return output, even if Gobuster fails
	if err != nil {
		return fmt.Sprintf("[Gobuster] Non-zero exit (%v):\n%s", err, string(output)), nil
	}

	return string(output), nil
}
