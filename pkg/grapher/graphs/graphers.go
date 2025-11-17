package graphs

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
)

type ApplyOptions interface {
	ApplyFlags(set *pflag.FlagSet) error
}

type Graphers[T any, U ApplyOptions] struct {
	graphs     map[string]func(defOpts *Options, opts U, data T) error
	applyFlags func(set *pflag.FlagSet)
	newOpts    func() U
}

func NewGraphers[T any, U ApplyOptions]() *Graphers[T, U] {
	return &Graphers[T, U]{
		graphs:     make(map[string]func(defOpts *Options, opts U, data T) error),
		applyFlags: AddFlags,
		newOpts:    nil,
	}
}

func (g *Graphers[T, U]) Process(data T, set *pflag.FlagSet) error {
	defOpts, err := DefaultOptions(set)
	if err != nil {
		return fmt.Errorf("failed to get default options: %w", err)
	}

	opts := g.newOpts()
	err = opts.ApplyFlags(set)
	if err != nil {
		return fmt.Errorf("failed to apply flags: %w", err)
	}

	log.Debug("Processing grapher", "defOpts", defOpts, "opts", opts)
	grapherFunc, ok := g.graphs[defOpts.GraphType]
	if !ok {
		graphs := strings.Join(slices.Collect(maps.Keys(g.graphs)), ",")

		return errors.New("unknown graph type: " + defOpts.GraphType + " (available: " + graphs + ")")
	}

	return grapherFunc(&defOpts, opts, data)
}

func (g *Graphers[T, U]) WithOptionsFunc(f func() U) *Graphers[T, U] {
	g.newOpts = f

	return g
}

// WithGrapher adds a new grapher function for the given graph type.
//	graphType: should be a registered graph type (e.g., "mermaid", "dot", etc.)
func (g *Graphers[T, U]) WithGrapher(graphType string, callback func(defOpts *Options, opts U, data T) error) *Graphers[T, U] {
	_, exists := g.graphs[graphType]
	if exists {
		panic("grapher for graph type " + graphType + " already exists")
	}
	g.graphs[graphType] = callback

	return g
}
func (g *Graphers[T, U]) WithApplyFlags(f func(set *pflag.FlagSet)) *Graphers[T, U] {
	old := g.applyFlags
	g.applyFlags = func(set *pflag.FlagSet) {
		old(set)
		f(set)
	}

	return g
}

func (g *Graphers[T, U]) AddFlags(set *pflag.FlagSet) {
	if g.applyFlags != nil {
		g.applyFlags(set)
	}
}
