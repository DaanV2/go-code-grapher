package golang

import (
	"os"
	"path/filepath"
	"strings"
)

// GetModulePath reads the module path from go.mod in the workspace root
func GetModulePath(filename string) (string, error) {
	filename = filepath.Clean(filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		l, found := strings.CutPrefix(strings.TrimSpace(line), "module ")
		if found {
			return l, nil
		}
	}

	return "", nil
}
