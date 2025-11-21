package markdown

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	// StartMarkerTemplate is the HTML comment that marks the start of an embedded section
	StartMarkerTemplate = "<!-- mermaid-embed-start:%s -->"
	// EndMarkerTemplate is the HTML comment that marks the end of an embedded section
	EndMarkerTemplate = "<!-- mermaid-embed-end:%s -->"
)

// EmbedSection represents a section in a markdown file that can be replaced
type EmbedSection struct {
	ID         string
	StartLine  int
	EndLine    int
	StartMarker string
	EndMarker   string
}

// FindEmbedSection finds the embed section with the given ID in the reader
func FindEmbedSection(r io.Reader, id string) (*EmbedSection, error) {
	startMarker := fmt.Sprintf(StartMarkerTemplate, id)
	endMarker := fmt.Sprintf(EndMarkerTemplate, id)
	
	scanner := bufio.NewScanner(r)
	lineNum := 0
	var section *EmbedSection
	
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		
		if strings.TrimSpace(line) == startMarker {
			if section != nil {
				return nil, fmt.Errorf("found nested start marker at line %d", lineNum)
			}
			section = &EmbedSection{
				ID:          id,
				StartLine:   lineNum,
				StartMarker: startMarker,
				EndMarker:   endMarker,
			}
		} else if strings.TrimSpace(line) == endMarker {
			if section == nil {
				return nil, fmt.Errorf("found end marker without start marker at line %d", lineNum)
			}
			section.EndLine = lineNum
			return section, nil
		}
	}
	
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	if section != nil {
		return nil, fmt.Errorf("found start marker but no end marker for ID '%s'", id)
	}
	
	return nil, fmt.Errorf("no embed section found with ID '%s'", id)
}

// ReplaceEmbedSection replaces the content of an embed section in a markdown file
func ReplaceEmbedSection(inputPath, id, newContent string) error {
	// Read the entire file
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	
	// Find the section
	section, err := FindEmbedSection(file, id)
	if err != nil {
		return err
	}
	
	// Read the file again to get all lines
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	var lines []string
	lineNum := 0
	
	for scanner.Scan() {
		lineNum++
		lines = append(lines, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	
	file.Close()
	
	// Build the new content
	var result []string
	
	// Add lines before the section (excluding the start marker line)
	if section.StartLine > 1 {
		result = append(result, lines[:section.StartLine-1]...)
	}
	
	// Add the new content (which includes the markers)
	contentLines := strings.Split(strings.TrimSpace(newContent), "\n")
	result = append(result, contentLines...)
	
	// Add lines after the section (excluding the end marker line)
	if section.EndLine < len(lines) {
		result = append(result, lines[section.EndLine:]...)
	}
	
	// Write the file back
	output := strings.Join(result, "\n")
	if !strings.HasSuffix(output, "\n") {
		output += "\n"
	}
	
	err = os.WriteFile(inputPath, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	return nil
}

// WrapWithMarkers wraps content with the appropriate start and end markers
func WrapWithMarkers(id, content string) string {
	startMarker := fmt.Sprintf(StartMarkerTemplate, id)
	endMarker := fmt.Sprintf(EndMarkerTemplate, id)
	
	// Remove any trailing newlines from content to have consistent formatting
	content = strings.TrimRight(content, "\n")
	
	return fmt.Sprintf("%s\n%s\n%s", startMarker, content, endMarker)
}

// ValidateID checks if an ID is valid (alphanumeric with hyphens and underscores)
func ValidateID(id string) error {
	if id == "" {
		return fmt.Errorf("ID cannot be empty")
	}
	
	matched, err := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, id)
	if err != nil {
		return err
	}
	
	if !matched {
		return fmt.Errorf("ID must contain only alphanumeric characters, hyphens, and underscores")
	}
	
	return nil
}
