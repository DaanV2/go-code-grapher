package mermaid

import (
	"errors"

	"github.com/daanv2/go-code-grapher/pkg/grapher/graphs"
	statediagrams "github.com/daanv2/go-code-grapher/pkg/grapher/graphs/state-diagrams"
)

type StateDiagramWriter struct {
	graphs.StringWriter
	opts  *statediagrams.Options
	ident string
}

func NewStateDiagramWriter(f graphs.StringCloserWriter, opts *statediagrams.Options) (statediagrams.Writer, error) {
	if f == nil || opts == nil {
		return nil, errors.New("file and options must be provided")
	}

	ident := "    "
	if opts.GraphOnly {
		ident = ""
	}

	return &StateDiagramWriter{
		graphs.StringWriter{
			Writer: f,
		},
		opts,
		ident,
	}, nil
}

func (s *StateDiagramWriter) Close() error {
	return s.StringWriter.Close()
}

// WriteState writes a state with the given id and description.
func (s *StateDiagramWriter) WriteState(id, description string) error {
	return s.Writef(s.ident+"%q : %q", id, description)
}

func (s *StateDiagramWriter) WriteTransition(fromID, toID, label string) error {
	if label == "" {
		return s.WriteLinef(s.ident+"%q --> %q", fromID, toID)
	}

	return s.WriteLinef(s.ident+"%q --> %q: %q", fromID, toID, label)
}

// Start implements statediagrams.Writer.
func (s *StateDiagramWriter) Start() error {
	if s.opts.GraphOnly {
		return nil
	}

	return s.WriteLines(
		"```mermaid",
		MakeMetadata(s.opts.Annotations),
		"stateDiagram-v2",
	)
}

// Finish implements statediagrams.Writer.
func (s *StateDiagramWriter) Finish() error {
	if s.opts.GraphOnly {
		return nil
	}

	return s.Write("\n```")
}
