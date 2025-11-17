package graphs

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/daanv2/go-code-grapher/pkg/extensions/xos"
	"github.com/spf13/pflag"
)

type Options struct {
	Filename  string
	Writer    StringCloserWriter
	GrapherID string
	GraphType string
}

func AddFlags(set *pflag.FlagSet) {
	set.StringToString("annotations", nil, "Annotations to add to the graph")
	set.String("output", "", "Output filename for the graph, if empty, output to stdout")

	graphers := AvailableGraphers()
	graphTypes := GetRegisteredTypes()

	set.String("grapher-id", graphers[0], "The grapher to use (options: "+strings.Join(graphers, ", ")+")")
	set.String("graph-type", graphTypes[0], "The type of graph to generate (options: "+strings.Join(graphTypes, ", ")+")")
}

func DefaultOptions(set *pflag.FlagSet) (Options, error) {
	opts := Options{}
	var err error

	opts.Filename, err = set.GetString("output")
	if err != nil {
		return opts, fmt.Errorf("failed to get output filename: %w", err)
	}
	opts.GrapherID, err = set.GetString("grapher-id")
	if err != nil {
		return opts, fmt.Errorf("failed to get grapher-id: %w", err)
	}
	opts.GraphType, err = set.GetString("graph-type")
	if err != nil {
		return opts, fmt.Errorf("failed to get graph-type: %w", err)
	}

	// Validate
	f, ok := GetFactory(opts.GraphType)
	if !ok {
		return opts, errors.New("invalid graph-type: " + opts.GraphType)
	}
	if !slices.Contains(f.GraphIds(), opts.GrapherID) {
		return opts, fmt.Errorf("invalid grapher-id: %s for graph-type: %s (options: %s)", opts.GrapherID, opts.GraphType, strings.Join(f.GraphIds(), ", "))
	}

	// Create writer
	opts.Writer, err = xos.OpenFileOr(opts.Filename)
	if err != nil {
		return opts, err
	}

	return opts, nil
}
