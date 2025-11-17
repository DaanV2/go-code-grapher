package statediagrams

import (
	"github.com/daanv2/go-code-grapher/pkg/grapher/graphs"
	"github.com/spf13/pflag"
)

type Options struct {
	graphs.BaseOptions
	GraphOnly bool // if true, only generate the graph without additional fluff around it (in case of mermaid, no markdown)
}

func NewOptions() *Options {
	return &Options{}
}

func AddFlags(set *pflag.FlagSet) {
	graphs.AddBaseOptionsFlags(set)

	set.Bool("graph-only", true, "If true, only generate the graph without additional fluff around it (in case of mermaid, no markdown)")
}

func (o *Options) ApplyFlags(set *pflag.FlagSet) error {
	err := o.BaseOptions.ApplyFlags(set)
	if err != nil {
		return err
	}

	graphOnly, err := set.GetBool("graph-only")
	if err != nil {
		return err
	}
	o.GraphOnly = graphOnly

	return nil
}
