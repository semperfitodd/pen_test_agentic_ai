package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"scanner/scanners"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./scanner <target>")
	}
	target := os.Args[1]

	allScanners := []scanners.Scanner{
		scanners.NewNmap(),
		scanners.NewAmass(),
		scanners.NewGobuster(),
		scanners.NewSQLMap(),
		scanners.NewNikto(),
		scanners.NewFfuf(),
	}

	var wg sync.WaitGroup
	results := make(chan scanners.ScanResult, len(allScanners))

	for _, scanner := range allScanners {
		wg.Add(1)
		go func(s scanners.Scanner) {
			defer wg.Done()
			output, err := s.Run(target)
			results <- scanners.ScanResult{
				Name:   s.Name(),
				Output: output,
				Error:  err,
			}
		}(scanner)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Printf("=== %s ===\n", result.Name)
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		} else {
			fmt.Println(result.Output)
		}
	}
}
