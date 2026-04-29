package extractor

import (
	"strings"
)

func RemoveComments(content string, fileType string) string {
	switch {
	case strings.HasSuffix(fileType, ".go"):
		return removeGoComments(content)
	case strings.HasSuffix(fileType, ".js"), strings.HasSuffix(fileType, ".ts"):
		return removeJSComments(content)
	case strings.HasSuffix(fileType, ".py"):
		return removePythonComments(content)
	case strings.HasSuffix(fileType, ".java"), strings.HasSuffix(fileType, ".c"), strings.HasSuffix(fileType, ".cpp"), strings.HasSuffix(fileType, ".h"):
		return removeCStyleComments(content)
	case strings.HasSuffix(fileType, ".yaml"), strings.HasSuffix(fileType, ".yml"):
		return removeYAMLComments(content)
	case strings.HasSuffix(fileType, ".toml"):
		return removeTOMLComments(content)
	case strings.HasSuffix(fileType, ".ini"):
		return removeINIComments(content)
	case strings.HasSuffix(fileType, ".rs"):
		return removeRustComments(content)
	case strings.HasSuffix(fileType, ".rb"), strings.HasSuffix(fileType, ".php"):
		return removeRubyPHPComments(content)
	default:
		return removeGenericComments(content)
	}
}

func removeGoComments(content string) string {
	content = removeBlockComments(content, "/*", "*/")

	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") {
			continue
		}

		newLine := removeInlineComment(line, "//")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeJSComments(content string) string {
	content = removeBlockComments(content, "/*", "*/")
	content = removeBlockComments(content, "<!--", "-->")

	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") {
			continue
		}

		newLine := removeInlineComment(line, "//")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removePythonComments(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		newLine := removeInlineComment(line, "#")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeCStyleComments(content string) string {
	content = removeBlockComments(content, "/*", "*/")

	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") {
			continue
		}

		newLine := removeInlineComment(line, "//")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeYAMLComments(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		newLine := removeInlineComment(line, "#")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeTOMLComments(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		newLine := removeInlineComment(line, "#")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeINIComments(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, ";") || strings.HasPrefix(trimmed, "#") {
			continue
		}

		newLine := removeInlineComment(line, ";")
		newLine = removeInlineComment(newLine, "#")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeRustComments(content string) string {
	content = removeBlockComments(content, "/*", "*/")

	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") {
			continue
		}

		newLine := removeInlineComment(line, "//")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeRubyPHPComments(content string) string {
	content = removeBlockComments(content, "/*", "*/")

	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		newLine := removeInlineComment(line, "#")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeGenericComments(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "//") {
			continue
		}

		newLine := removeInlineComment(line, "#")
		newLine = removeInlineComment(newLine, "//")
		result = append(result, newLine)
	}

	return strings.Join(result, "\n")
}

func removeBlockComments(content, start, end string) string {
	var result strings.Builder
	i := 0

	for i < len(content) {
		if i+len(start) <= len(content) && content[i:i+len(start)] == start {
			endIndex := strings.Index(content[i+len(start):], end)
			if endIndex != -1 {
				i += len(start) + endIndex + len(end)
				continue
			}
		}
		result.WriteByte(content[i])
		i++
	}

	return result.String()
}

func removeInlineComment(line, commentMarker string) string {
	inString := false
	stringChar := rune(0)

	for i := 0; i < len(line); i++ {
		char := line[i]

		if !inString && (char == '"' || char == '\'' || char == '`') {
			inString = true
			stringChar = rune(char)
		} else if inString && rune(char) == stringChar {
			inString = false
		}

		if !inString && i+len(commentMarker) <= len(line) && line[i:i+len(commentMarker)] == commentMarker {
			return strings.TrimSpace(line[:i])
		}
	}

	return line
}