package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/crystian/llm-context-extractor/extractor"
)

func main() {
	inputDir := flag.String("input", "", "Directory to extract code from (required)")
	outputFile := flag.String("output", "context.json", "Output JSON file path")
	verbose := flag.Bool("verbose", false, "Show detailed progress")
	excludePatterns := flag.String("exclude", "", "Additional patterns to exclude (comma-separated)")
	flag.Parse()

	if *inputDir == "" {
		fmt.Println("Error: -input flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(*inputDir); os.IsNotExist(err) {
		fmt.Printf("Error: input directory does not exist: %s\n", *inputDir)
		os.Exit(1)
	}

	absInputDir, err := filepath.Abs(*inputDir)
	if err != nil {
		fmt.Printf("Error: cannot get absolute path for input directory: %v\n", err)
		os.Exit(1)
	}

	// Parse exclude patterns
	var customExcludePatterns []string
	if *excludePatterns != "" {
		patterns := strings.Split(*excludePatterns, ",")
		for _, pattern := range patterns {
			pattern = strings.TrimSpace(pattern)
			if pattern != "" {
				customExcludePatterns = append(customExcludePatterns, pattern)
			}
		}
	}

	e := extractor.NewExtractorWithExcludes(*verbose, customExcludePatterns)

	if len(customExcludePatterns) > 0 && *verbose {
		fmt.Printf("Custom exclude patterns: %v\n", customExcludePatterns)
	}

	fmt.Printf("Extracting code from: %s\n", absInputDir)
	fmt.Printf("Output file: %s\n", *outputFile)

	if err := e.ExtractToJSON(absInputDir, *outputFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully extracted code to %s\n", *outputFile)
}