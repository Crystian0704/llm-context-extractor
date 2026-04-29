package extractor

import (
	"testing"
)

func TestShouldInclude(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"Go source file", "main.go", true},
		{"JavaScript file", "app.js", true},
		{"YAML config", "config.yaml", true},
		{"Dockerfile", "Dockerfile", true},
		{"Makefile", "Makefile", true},
		{"README", "README.md", false},
		{"Test file", "main_test.go", false},
		{"Text file", "notes.txt", false},
		{"Git directory", ".git", false},
		{"Node modules", "node_modules", false},
		{"Vendor directory", "vendor", false},
		{"Test directory", "test", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ShouldInclude(tt.path)
			if result != tt.expected {
				t.Errorf("ShouldInclude(%q) = %v, want %v", tt.path, result, tt.expected)
			}
		})
	}
}

func TestRemoveComments(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		fileType string
		expected string
	}{
		{
			name:     "Go single line comment",
			content:  "x := 42 // inline comment",
			fileType: "main.go",
			expected: "x := 42",
		},
		{
			name:     "Go multi-line comment",
			content:  "/* comment */\nx := 42",
			fileType: "main.go",
			expected: "\nx := 42",
		},
		{
			name:     "YAML comment",
			content:  "port: 8080 # server port",
			fileType: "config.yaml",
			expected: "port: 8080",
		},
		{
			name:     "Python comment",
			content:  "x = 42 # inline comment",
			fileType: "app.py",
			expected: "x = 42",
		},
		{
			name:     "Preserve string with comment marker",
			content:  `url := "http://example.com"`,
			fileType: "main.go",
			expected: `url := "http://example.com"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveComments(tt.content, tt.fileType)
			if result != tt.expected {
				t.Errorf("RemoveComments(%q, %q) = %q, want %q", tt.content, tt.fileType, result, tt.expected)
			}
		})
	}
}