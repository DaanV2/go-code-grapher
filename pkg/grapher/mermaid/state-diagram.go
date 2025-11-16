package mermaid

import (
	"fmt"
	"os"
	"path/filepath"
)

type StateDiagramWriter struct {
	w *os.File
}

func NewStateDiagramWriter(filename string) (*StateDiagramWriter, error) {
	filename = filepath.Clean(filename)
	w, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return &StateDiagramWriter{
		w: w,
	}, nil
}

func (s *StateDiagramWriter) Close() error {
	return s.w.Close()
}

func (s *StateDiagramWriter) Write(content string) error {
	_, err := s.w.WriteString(content)

	return err
}

// WriteState writes a state with the given id and description.
func (s *StateDiagramWriter) WriteState(id, description string) error {
	return s.Write(fmt.Sprintf("    %q : %q", id, description))
}

func (s *StateDiagramWriter) WriteTransition(fromID, toID, label string) error {
	if label == "" {
		return s.Write(fmt.Sprintf("    %q --> %q", fromID, toID))
	}

	return s.Write(fmt.Sprintf("    %q --> %q: %q", fromID, toID, label))
}