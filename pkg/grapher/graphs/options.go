package graphs

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/daanv2/go-code-grapher/pkg/extensions/xos"
	"github.com/daanv2/go-code-grapher/pkg/markdown"
	"github.com/spf13/pflag"
)

type Options struct {
	Filename           string
	Writer             StringCloserWriter
	GrapherID          string
	GraphType          string
	MarkdownEmbedInto  string
	MarkdownEmbedID    string
	embedBuffer        *bytes.Buffer
}

func AddFlags(set *pflag.FlagSet) {
	set.String("output", "", "Output filename for the graph, if empty, output to stdout")

	graphers := AvailableGraphers()
	graphTypes := GetRegisteredTypes()

	set.String("grapher-id", graphers[0], "The grapher to use (options: "+strings.Join(graphers, ", ")+")")
	set.String("graph-type", graphTypes[0], "The type of graph to generate (options: "+strings.Join(graphTypes, ", ")+")")
	
	// Markdown embedding flags
	set.String("markdown-embed-into", "", "Markdown file to embed the generated graph into")
	set.String("markdown-embed-id", "", "ID of the section in the markdown file to replace")
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
	
	// Get markdown embedding options
	opts.MarkdownEmbedInto, err = set.GetString("markdown-embed-into")
	if err != nil {
		return opts, fmt.Errorf("failed to get markdown-embed-into: %w", err)
	}
	opts.MarkdownEmbedID, err = set.GetString("markdown-embed-id")
	if err != nil {
		return opts, fmt.Errorf("failed to get markdown-embed-id: %w", err)
	}

	// Validate
	f, ok := GetFactory(opts.GraphType)
	if !ok {
		return opts, errors.New("invalid graph-type: " + opts.GraphType)
	}
	if !slices.Contains(f.GraphIds(), opts.GrapherID) {
		return opts, fmt.Errorf("invalid grapher-id: %s for graph-type: %s (options: %s)", opts.GrapherID, opts.GraphType, strings.Join(f.GraphIds(), ", "))
	}
	
	// Validate markdown embedding options
	if opts.MarkdownEmbedInto != "" {
		if opts.MarkdownEmbedID == "" {
			return opts, errors.New("markdown-embed-id is required when markdown-embed-into is specified")
		}
		if err := markdown.ValidateID(opts.MarkdownEmbedID); err != nil {
			return opts, fmt.Errorf("invalid markdown-embed-id: %w", err)
		}
		
		// When embedding, write to a buffer first
		opts.embedBuffer = &bytes.Buffer{}
		opts.Writer = &bufferCloserWriter{Buffer: opts.embedBuffer}
	} else {
		// Create writer for normal output
		opts.Writer, err = xos.OpenFileOr(opts.Filename)
		if err != nil {
			return opts, err
		}
	}

	return opts, nil
}

// bufferCloserWriter wraps a bytes.Buffer to implement StringCloserWriter
type bufferCloserWriter struct {
	*bytes.Buffer
}

func (b *bufferCloserWriter) Close() error {
	return nil
}

// FinalizeMarkdownEmbed embeds the generated graph into the markdown file
func (opts *Options) FinalizeMarkdownEmbed() error {
	if opts.MarkdownEmbedInto == "" {
		return nil
	}
	
	if opts.embedBuffer == nil {
		return errors.New("embed buffer is nil")
	}
	
	// Get the generated content
	content := opts.embedBuffer.String()
	
	// Wrap it with markers
	wrappedContent := markdown.WrapWithMarkers(opts.MarkdownEmbedID, content)
	
	// Replace the section in the markdown file
	return markdown.ReplaceEmbedSection(opts.MarkdownEmbedInto, opts.MarkdownEmbedID, wrappedContent)
}
