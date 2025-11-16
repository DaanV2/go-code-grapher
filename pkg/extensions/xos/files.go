package xos

import (
	"iter"
	"os"
	"path/filepath"
)

func GetFiles(dir string) ([]string, error) {
	files := make([]string, 0, 10)
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

func AllFiles(dirs []string, recursive bool) iter.Seq[string] {
	return func(yield func(string) bool) {
		todo := make([]string, 0, len(dirs))
		todo = append(todo, dirs...)

		for i := 0; i < len(todo); i++ {
			dir := todo[i]
			entries, err := os.ReadDir(dir)
			if err != nil {
				continue
			}
			for _, entry := range entries {
				p := filepath.Join(dir, entry.Name())
				if entry.IsDir() {
					if recursive {
						todo = append(todo, p)
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

func AllGoFiles(dirs []string, recursive bool) iter.Seq[string] {
	return func(yield func(string) bool) {
		for file := range AllFiles(dirs, recursive) {
			if filepath.Ext(file) == ".go" {
				if !yield(file) {
					return
				}
			}
		}
	}
}
