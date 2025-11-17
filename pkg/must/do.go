package must

import (
	"fmt"
	"runtime"
)

func Do(err error) {
	if err != nil {
		// Append stack
		_, file, line, ok := runtime.Caller(1)
		if ok {
			err = fmt.Errorf("%w\n,  %s:%d", err, file, line)
		}

		panic(err)
	}
}
