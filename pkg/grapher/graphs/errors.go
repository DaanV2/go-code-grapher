package graphs

import (
	"fmt"
	"iter"
	"slices"
)

func ErrUnknownOption[T ~string](message string, options iter.Seq[T]) error {
	opts := slices.Collect(options)

	return fmt.Errorf(message, opts)
}