package graphs

import (
	"maps"
	"slices"

	"github.com/daanv2/go-code-grapher/pkg/extensions/xslices"
)

var factories = make(map[string]RegisterFactory)

type RegisterFactory interface {
	GraphIds() []string
}

func Register[TGraph BaseWriter, TOption any](diagramType string) *DiagramFactory[TGraph, TOption] {
	factory := NewDiagramFactory[TGraph, TOption](diagramType)
	factories[factory.Type] = factory

	return factory
}

// GetRegisteredTypes returns the registered diagram types.
func GetRegisteredTypes() []string { return slices.Collect(maps.Keys(factories)) }

func GetFactory(diagramType string) (RegisterFactory, bool) {
	factory, ok := factories[diagramType]

	return factory, ok
}

// AvailableGraphers returns all available grapher IDs across all registered diagram types.
func AvailableGraphers() []string {
	var result []string

	for _, graphT := range GetRegisteredTypes() {
		f, _ := GetFactory(graphT)
		result = append(result, f.GraphIds()...)
	}

	return xslices.Unique(result)
}