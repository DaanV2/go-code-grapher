package xflags

import (
	"fmt"

	"github.com/spf13/pflag"
)

func SetIfUnchanged(set *pflag.FlagSet, name, value string) error {
	flag := set.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flag %s not found", name)
	}
	if !flag.Changed {
		return flag.Value.Set(value)
	}

	return nil
}
