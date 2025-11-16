package xslices

import "slices"

// Unique returns a new slice containing only the unique elements from the input slice.
func Unique[S ~[]E, E comparable](slice S) S {
	if len(slice) == 0 {
		return slice
	}
	// For small slices, a simple O(n^2) algorithm is faster due to lower overhead.
	if len(slice) < 100 {
		result := make(S, 0, len(slice))

		for _, v := range slice {
			if !slices.Contains(slice, v) {
				result = append(result, v)
			}
		}

		return result
	}

	s := make(map[E]struct{}, len(slice))
	result := make(S, 0, len(s))
	for _, v := range slice {
		if _, exists := s[v]; !exists {
			s[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}