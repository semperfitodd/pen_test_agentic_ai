package scanners

import (
	"os/exec"
)

type Ffuf struct{}

func NewFfuf() *Ffuf { return &Ffuf{} }

func (f *Ffuf) Name() string { return "ffuf" }

func (f *Ffuf) Run(target string) (string, error) {
	cmd := exec.Command(
	"ffuf",
	"-u", "https://"+target+"/FUZZ",
	"-w", "/usr/share/wordlists/dirb/common.txt",
	"-mc", "200,204,301,302,403",
	"-r",
	"-k",
	"-timeout", "10",
    )
	output, err := cmd.CombinedOutput()
	return string(output), err
}
