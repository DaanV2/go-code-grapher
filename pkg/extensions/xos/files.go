package xos

import (
	"iter"
	"os"
	"path/filepath"
)

func GetFiles(dir string) ([]string, error) {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		files = append(files, filepath.Join(dir, entry.Name()))
	}
	return files, nil
}

func AllFiles(dirs []string, recurse bool) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, dir := range dirs {
			entries, err := os.ReadDir(dir)
			if err != nil {
				continue
			}
			for _, entry := range entries {
				p := filepath.Join(dir, entry.Name())
				if entry.IsDir() {
					if recurse {
						dirs = append(dirs, p)
					}
				} else {
					if !yield(p) {
						return
					}
				}
			}
		}
	}
}

func AllGoFiles(dirs []string, recurse bool) iter.Seq[string] {
	return func(yield func(string) bool) {
		for file := range AllFiles(dirs, recurse) {
			if filepath.Ext(file) == ".go" {
				if !yield(file) {
					return
				}
			}
		}
	}
}
