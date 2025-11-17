package grapher

import (
	"fmt"

	"github.com/daanv2/go-code-grapher/pkg/ast"
	"github.com/daanv2/go-code-grapher/pkg/grapher/graphs"
	statediagrams "github.com/daanv2/go-code-grapher/pkg/grapher/graphs/state-diagrams"
)

var ImportsGraphers = graphs.NewGraphers[*ast.ImportCollector, *statediagrams.Options]().
	WithOptionsFunc(statediagrams.NewOptions).
	WithApplyFlags(statediagrams.AddFlags).
	WithGrapher("state-diagram", func(defOpts *graphs.Options, opts *statediagrams.Options, col *ast.ImportCollector) error {
		grapher, err := StateDiagramFactory.Create(defOpts.Writer, defOpts.GrapherID, opts)
		if err != nil {
			return fmt.Errorf("failed to create state diagram grapher: %w", err)
		}

		for pack, imports := range col.Imports() {
			for _, imp := range imports {
				err := grapher.WriteTransition(pack, imp, "")
				if err != nil {
					return fmt.Errorf("failed to write transition from %s to %s: %w", pack, imp, err)
				}
			}
		}

		return nil
	})
