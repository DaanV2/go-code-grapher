package xos

import (
	"os"
	"path/filepath"
)

// OpenFileOr opens the file with the given filename for writing.
// If the filename is empty, it returns os.Stdout.
func OpenFileOr(filename string) (*os.File, error) {
	if filename == "" {
		return os.Stdout, nil
	}

	// Create or truncate the file
	return os.Create(filepath.Clean(filename))
}