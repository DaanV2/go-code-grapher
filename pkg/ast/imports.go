package ast

import (
	"path/filepath"
	"strings"

	"github.com/daanv2/go-code-grapher/pkg/extensions/xslices"
	"github.com/daanv2/go-code-grapher/pkg/golang"
)

type ImportCollector struct {
	// Package -> Imported Packages
	imports map[string][]string

	// Dir -> Package mapping
	dirPackages map[string][]string

	moduleName string // "github.com/<org>/<repo>"
	moduleDir  string // Absolute path to module directory
}

func NewImportCollector(modulePath string) (*ImportCollector, error) {
	modulePath, err := filepath.Abs(modulePath)
	if err != nil {
		return nil, err
	}

	// Get module path from go.mod
	moduleName, err := golang.GetModulePath(modulePath)
	if err != nil {
		return nil, err
	}

	return &ImportCollector{
		imports: make(map[string][]string),
		dirPackages: make(map[string][]string),
		moduleName: moduleName,
		moduleDir:  filepath.Dir(modulePath),
	}, nil
}

// Collect parses the given Go file and collects its imports.
func (c *ImportCollector) Collect(filename string) error {
	f, err := ParseFile(filename)
	if err != nil {
		return err
	}

	// Fully qualified package name
	pack, err := golang.QualifyPackageName(c.moduleDir, c.moduleName, filepath.Dir(filename), f.Name.Name)
	if err != nil {
		return err
	}

	imps := make([]string, 0, len(f.Imports))
	for _, i := range f.Imports {
		v := strings.Trim(i.Path.Value, "\\\"")
		imps = append(imps, v)
	}

	c.imports[pack] = xslices.Unique(append(c.imports[pack], imps...))
	dir := filepath.Dir(filename)

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