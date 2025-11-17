package graphs

import (
	"maps"

	"github.com/spf13/pflag"
)

type BaseOptions struct {
	Annotations map[string]string
}

func AddBaseOptionsFlags(set *pflag.FlagSet) {
	set.StringToString("annotations", nil, "Annotations to add to the graph")
}

func (o *BaseOptions) ApplyFlags(set *pflag.FlagSet) error {
	annotations, err := set.GetStringToString("annotations")
	if err != nil {
		return err
	}

	if o.Annotations == nil {
		o.Annotations = make(map[string]string)
	} else {
		maps.Copy(o.Annotations, annotations)
	}

	return nil
}