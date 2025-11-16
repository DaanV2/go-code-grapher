package ast

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/charmbracelet/log"
)

func ParseFile(filename string) (*ast.File, error) {
	log.Debug("parsing filename", "filename", filename)

	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	return f, nil
}
