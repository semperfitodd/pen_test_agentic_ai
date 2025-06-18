package scanners

import (
	"os/exec"
)

type SQLMap struct{}

func NewSQLMap() *SQLMap { return &SQLMap{} }

func (s *SQLMap) Name() string { return "SQLMap" }

func (s *SQLMap) Run(target string) (string, error) {
	cmd := exec.Command(
	"sqlmap",
	"-u", "https://"+target,
	"--batch",
	"--random-agent",
	"--level=5",
	"--risk=3",
	"--crawl=3",
	"--threads=10",
    )
	output, err := cmd.CombinedOutput()
	return string(output), err
}
