package ast

import "github.com/daanv2/go-code-grapher/pkg/extensions/xslices"

type ImportCollector struct {
	// Package -> Imported Packages
	imports map[string][]string
}

func NewImportCollector() *ImportCollector {
	return &ImportCollector{
		imports: make(map[string][]string),
	}
}

func (c *ImportCollector) Collect(filePath string) error {
	f, err := ParseFile(filePath)
	if err != nil {
		return err
	}

	pack := f.Name.Name
	imps := make([]string, 0, len(f.Imports))
	for _, i := range f.Imports {
		imps = append(imps, i.Path.Value)
	}

	c.imports[pack] = xslices.Unique(append(c.imports[pack], imps...))

	return nil
}
