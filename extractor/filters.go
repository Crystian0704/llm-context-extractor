package extractor

import (
	"path/filepath"
	"strings"
)

var (
	sourceExtensions = []string{".go", ".js", ".ts", ".py", ".java", ".c", ".cpp", ".h", ".rs", ".rb", ".php"}
	configExtensions = []string{".yaml", ".yml", ".json", ".toml", ".ini", ".cfg"}
	infraFiles       = []string{"Dockerfile", "dockerfile", "Makefile", "makefile"}

	excludedDirs      = []string{".git", "node_modules", "vendor", "test", "spec", "tests", "__tests__", ".venv", "venv", "env", ".env", "dist", "build", "target", "bin", "obj"}
	excludedFiles     = []string{"README", "CHANGELOG", "LICENSE", "CONTRIBUTING"}
	excludedExtensions = []string{".md", ".txt", ".rst", ".log"}
	testPatterns      = []string{"_test.go", ".test.", ".spec."}
)

func ShouldInclude(path string) bool {
	base := filepath.Base(path)
	ext := strings.ToLower(filepath.Ext(path))

	if IsExcludedDir(base) {
		return false
	}

	if IsExcludedFile(base) {
		return false
	}

	if contains(excludedExtensions, ext) {
		return false
	}

	if isTestFile(base) {
		return false
	}

	return IsSourceFile(ext) || IsConfigFile(ext) || IsInfraFile(base)
}

func IsExcludedDir(dir string) bool {
	return contains(excludedDirs, dir)
}

func IsExcludedFile(file string) bool {
	base := strings.ToLower(file)
	for _, excluded := range excludedFiles {
		if strings.HasPrefix(base, strings.ToLower(excluded)) {
			return true
		}
	}
	return false
}

func IsSourceFile(ext string) bool {
	return contains(sourceExtensions, ext)
}

func IsConfigFile(ext string) bool {
	return contains(configExtensions, ext)
}

func IsInfraFile(file string) bool {
	return contains(infraFiles, file)
}

func isTestFile(file string) bool {
	lowerFile := strings.ToLower(file)
	for _, pattern := range testPatterns {
		if strings.Contains(lowerFile, strings.ToLower(pattern)) {
			return true
		}
	}
	return false
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}