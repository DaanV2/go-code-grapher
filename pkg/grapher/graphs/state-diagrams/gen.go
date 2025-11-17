package statediagrams

import "github.com/daanv2/go-code-grapher/pkg/grapher/graphs"

type Writer interface {
	graphs.BaseWriter
	WriteState(id, description string) error
	WriteTransition(fromID, toID, label string) error
}
