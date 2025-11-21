package markdown_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/daanv2/go-code-grapher/pkg/markdown"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{"valid simple", "test", false},
		{"valid with hyphen", "test-id", false},
		{"valid with underscore", "test_id", false},
		{"valid with numbers", "test123", false},
		{"valid complex", "my-test_id-123", false},
		{"empty", "", true},
		{"with spaces", "test id", true},
		{"with special chars", "test@id", true},
		{"with dots", "test.id", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := markdown.ValidateID(tt.id)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWrapWithMarkers(t *testing.T) {
	id := "test-diagram"
	content := "```mermaid\ngraph TD\n    A --> B\n```"
	
	result := markdown.WrapWithMarkers(id, content)
	
	assert.Contains(t, result, "<!-- mermaid-embed-start:test-diagram -->")
	assert.Contains(t, result, "<!-- mermaid-embed-end:test-diagram -->")
	assert.Contains(t, result, content)
}

func TestFindEmbedSection(t *testing.T) {
	tests := []struct {
		name    string
		content string
		id      string
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid section",
			content: `# Test
<!-- mermaid-embed-start:test -->
content here
<!-- mermaid-embed-end:test -->
more content`,
			id:      "test",
			wantErr: false,
		},
		{
			name: "section not found",
			content: `# Test
<!-- mermaid-embed-start:other -->
content here
<!-- mermaid-embed-end:other -->`,
			id:      "test",
			wantErr: true,
			errMsg:  "no embed section found",
		},
		{
			name: "missing end marker",
			content: `# Test
<!-- mermaid-embed-start:test -->
content here`,
			id:      "test",
			wantErr: true,
			errMsg:  "no end marker",
		},
		{
			name: "missing start marker",
			content: `# Test
content here
<!-- mermaid-embed-end:test -->`,
			id:      "test",
			wantErr: true,
			errMsg:  "end marker without start marker",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.content)
			section, err := markdown.FindEmbedSection(reader, tt.id)
			
			if tt.wantErr {
				require.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				require.NoError(t, err)
				assert.NotNil(t, section)
				assert.Equal(t, tt.id, section.ID)
				assert.Greater(t, section.EndLine, section.StartLine)
			}
		})
	}
}

func TestReplaceEmbedSection(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	
	testFile := filepath.Join(tmpDir, "test.md")
	originalContent := `# Test Document

Some text before.

<!-- mermaid-embed-start:architecture -->
` + "```mermaid" + `
graph TD
    Old --> Content
` + "```" + `
<!-- mermaid-embed-end:architecture -->

Some text after.
`
	
	err := os.WriteFile(testFile, []byte(originalContent), 0600)
	require.NoError(t, err)
	
	newContent := markdown.WrapWithMarkers("architecture", "```mermaid\ngraph TD\n    New --> Content\n```")
	
	err = markdown.ReplaceEmbedSection(testFile, "architecture", newContent)
	require.NoError(t, err)
	
	// Read the file back
	result, err := os.ReadFile(testFile) // #nosec G304 -- test file path is controlled
	require.NoError(t, err)
	
	resultStr := string(result)
	
	// Verify the content was replaced
	assert.Contains(t, resultStr, "New --> Content")
	assert.NotContains(t, resultStr, "Old --> Content")
	assert.Contains(t, resultStr, "Some text before.")
	assert.Contains(t, resultStr, "Some text after.")
	assert.Contains(t, resultStr, "<!-- mermaid-embed-start:architecture -->")
	assert.Contains(t, resultStr, "<!-- mermaid-embed-end:architecture -->")
}

func TestReplaceEmbedSection_NonExistentFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "nonexistent.md")
	
	newContent := markdown.WrapWithMarkers("test", "content")
	err := markdown.ReplaceEmbedSection(testFile, "test", newContent)
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open file")
}

func TestReplaceEmbedSection_NonExistentSection(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	
	originalContent := `# Test
No markers here
`
	err := os.WriteFile(testFile, []byte(originalContent), 0600)
	require.NoError(t, err)
	
	newContent := markdown.WrapWithMarkers("test", "content")
	err = markdown.ReplaceEmbedSection(testFile, "test", newContent)
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no embed section found")
}
