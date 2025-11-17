package mermaid

import (
	"strings"
)

func MakeMetadata(data map[string]string) string {
	b := strings.Builder{}

	b.WriteString("---\n")
	for key, value := range data {
		_, _ = b.WriteString(key)
		_, _ = b.WriteString(": ")
		_, _ = b.WriteString(value)
		_, _ = b.WriteString("\n")
	}
	b.WriteString("---\n")

	return b.String()
}