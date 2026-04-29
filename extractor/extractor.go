package extractor

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Extractor struct {
	Verbose         bool
	ExcludePatterns []string
}

func NewExtractor(verbose bool) *Extractor {
	return &Extractor{
		Verbose:         verbose,
		ExcludePatterns: []string{},
	}
}

func NewExtractorWithExcludes(verbose bool, excludePatterns []string) *Extractor {
	return &Extractor{
		Verbose:         verbose,
		ExcludePatterns: excludePatterns,
	}
}

func (e *Extractor) Extract(dir string) (map[string]string, error) {
	result := make(map[string]string)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		// Check custom exclude patterns first
		if e.matchesExcludePattern(relPath, d.IsDir()) {
			if e.Verbose {
				if d.IsDir() {
					fmt.Printf("Excluding directory (custom pattern): %s\n", relPath)
				} else {
					fmt.Printf("Skipping file (custom pattern): %s\n", relPath)
				}
			}
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.IsDir() {
			if IsExcludedDir(d.Name()) {
				if e.Verbose {
					fmt.Printf("Excluding directory: %s\n", relPath)
				}
				return filepath.SkipDir
			}
			return nil
		}

		if !ShouldInclude(path) {
			if e.Verbose {
				fmt.Printf("Skipping file: %s\n", relPath)
			}
			return nil
		}

		if e.Verbose {
			fmt.Printf("Processing file: %s\n", relPath)
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", path, err)
		}

		cleanContent := RemoveComments(string(content), filepath.Base(path))

		result[relPath] = cleanContent

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	return result, nil
}

func (e *Extractor) matchesExcludePattern(path string, isDir bool) bool {
	if len(e.ExcludePatterns) == 0 {
		return false
	}

	// Convert path to use forward slashes for pattern matching
	normalizedPath := filepath.ToSlash(path)

	for _, pattern := range e.ExcludePatterns {
		normalizedPattern := filepath.ToSlash(pattern)

		// Handle patterns ending with /** (exclude entire directory contents)
		if strings.HasSuffix(normalizedPattern, "/**") {
			dirPattern := strings.TrimSuffix(normalizedPattern, "/**")
			if strings.HasPrefix(normalizedPath, dirPattern+"/") || normalizedPath == dirPattern {
				return true
			}
			continue
		}

		// Handle patterns ending with /*/ (exclude immediate subdirectories)
		if strings.HasSuffix(normalizedPattern, "/*/") {
			dirPattern := strings.TrimSuffix(normalizedPattern, "/*/")
			// Check if path is an immediate subdirectory
			pathParts := strings.Split(normalizedPath, "/")
			if len(pathParts) == 2 && pathParts[0] == dirPattern {
				return true
			}
			// Check if path is inside an immediate subdirectory
			if len(pathParts) >= 2 && pathParts[0] == dirPattern {
				return true
			}
			continue
		}

		// Check for exact match
		if normalizedPath == normalizedPattern {
			return true
		}

		// Check if path starts with pattern (for directories)
		if strings.HasPrefix(normalizedPath, normalizedPattern+"/") {
			return true
		}

		// Check if path ends with pattern (for files)
		if strings.HasSuffix(normalizedPath, normalizedPattern) {
			return true
		}

		// Check for wildcard patterns
		if strings.Contains(normalizedPattern, "*") {
			matched, err := filepath.Match(normalizedPattern, filepath.Base(path))
			if err == nil && matched {
				return true
			}

			// Try matching against full path
			matched, err = filepath.Match(normalizedPattern, normalizedPath)
			if err == nil && matched {
				return true
			}

			// Try matching path components
			pathParts := strings.Split(normalizedPath, "/")
			for _, part := range pathParts {
				matched, err := filepath.Match(normalizedPattern, part)
				if err == nil && matched {
					return true
				}
			}
		}

		// Check if pattern is a directory and path is inside it
		if isDir && strings.HasPrefix(normalizedPattern, normalizedPath+"/") {
			return true
		}
	}

	return false
}

func (e *Extractor) ExtractToJSON(dir string, outputPath string) error {
	result, err := e.Extract(dir)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	if e.Verbose {
		fmt.Printf("Successfully extracted %d files to %s\n", len(result), outputPath)
	}

	return nil
}

func Extract(dir string) (map[string]string, error) {
	e := NewExtractor(false)
	return e.Extract(dir)
}

func ExtractToJSON(dir string, outputPath string) error {
	e := NewExtractor(false)
	return e.ExtractToJSON(dir, outputPath)
}