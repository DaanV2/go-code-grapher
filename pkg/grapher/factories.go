package grapher

import (
	"github.com/daanv2/go-code-grapher/pkg/grapher/graphs"
	statediagrams "github.com/daanv2/go-code-grapher/pkg/grapher/graphs/state-diagrams"
	"github.com/daanv2/go-code-grapher/pkg/grapher/mermaid"
)

var (
	StateDiagramFactory = graphs.Register[statediagrams.Writer, statediagrams.Options]("state-diagram").
		With(mermaid.GRAPH_ID, mermaid.NewStateDiagramWriter)

)

