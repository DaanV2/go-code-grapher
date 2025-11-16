package ast

import (
	"path/filepath"
	"strings"

	"github.com/daanv2/go-code-grapher/pkg/extensions/xslices"
)

type ImportCollector struct {
	// Package -> Imported Packages
	imports map[string][]string

	// Dir -> Package mapping
	dirPackages map[string][]string
}

func NewImportCollector() *ImportCollector {
	return &ImportCollector{
		imports: make(map[string][]string),
		dirPackages: make(map[string][]string),
	}
}

// Collect parses the given Go file and collects its imports.
func (c *ImportCollector) Collect(filePath string) error {
	f, err := ParseFile(filePath)
	if err != nil {
		return err
	}

	pack := f.Name.Name
	imps := make([]string, 0, len(f.Imports))
	for _, i := range f.Imports {
		v := strings.Trim(i.Path.Value, "\\\"")
		imps = append(imps, v)
	}

	c.imports[pack] = xslices.Unique(append(c.imports[pack], imps...))
	dir := filepath.Dir(filePath)

	c.dirPackages[dir] = xslices.Unique(append(c.dirPackages[dir], pack))

	return nil
}

// Imports returns the collected imports mapping.
func (c *ImportCollector) Imports() map[string][]string {
	return c.imports
}

// DirPackages returns a mapping of directories to the packages they contain.
func (c *ImportCollector) DirPackages() map[string][]string {
	return c.dirPackages
}