package graphs

import (
	"maps"
	"slices"
)

type DiagramConstructor[TGraph BaseWriter, TOption any] func(f StringCloserWriter, opts *TOption) (TGraph, error)

type DiagramFactory[TGraph BaseWriter, TOption any] struct {
	Type         string
	Constructors map[string]DiagramConstructor[TGraph, TOption]
}

func NewDiagramFactory[TGraph BaseWriter, TOption any](diagramType string) *DiagramFactory[TGraph, TOption] {
	return &DiagramFactory[TGraph, TOption]{
		Type:         diagramType,
		Constructors: make(map[string]DiagramConstructor[TGraph, TOption]),
	}
}

func (fact *DiagramFactory[T, U]) Get(grapherId string) (DiagramConstructor[T, U], bool) {
	constructor, ok := fact.Constructors[grapherId]

	return constructor, ok
}

func (fact *DiagramFactory[T, U]) GraphIds() []string {
	return slices.Collect(maps.Keys(fact.Constructors))
}

func (fact *DiagramFactory[T, U]) Create(f StringCloserWriter, grapherId string, opts *U) (T, error) {
	var result T
	constructor, ok := fact.Get(grapherId)
	if ok {
		return constructor(f, opts)
	}

	return result, ErrUnknownOption("unknown graphid for: "+fact.Type+" (options: %v)", maps.Keys(fact.Constructors))
}

func (fact *DiagramFactory[T, U]) With(grapherId string, constructor DiagramConstructor[T, U]) *DiagramFactory[T, U] {
	fact.Constructors[grapherId] = constructor

	return fact
}
