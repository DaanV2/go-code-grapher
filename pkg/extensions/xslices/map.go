package xslices

func Map[S any, T any](source []S, mapper func(S) T) []T {
	result := make([]T, len(source))
	for i, v := range source {
		result[i] = mapper(v)
	}

	return result
}

func MapE[S any, T any](source []S, mapper func(S) (T, error)) ([]T, error) {
	result := make([]T, len(source))
	var err error
	for i, v := range source {
		result[i], err = mapper(v)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
