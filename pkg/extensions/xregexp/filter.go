package xregexp

import (
	"fmt"
	"regexp"
)

// Filter represents a regex-based filter for strings.
type Filter struct {
	patterns []*regexp.Regexp
}

func FromPatterns(patterns []string) (*Filter, error) {
	f := &Filter{}
	err := f.AddPatterns(patterns...)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *Filter) Len() int {
	return len(f.patterns)
}

// AddPatterns adds multiple regex patterns to the filter.
func (f *Filter) AddPatterns(patterns ...string) error {
	for _, pattern := range patterns {
		err := f.AddPattern(pattern)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddPattern adds a single regex pattern to the filter.
func (f *Filter) AddPattern(pattern string) error {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("failed to compile regex pattern %q: %w", pattern, err)
	}
	f.patterns = append(f.patterns, regex)

	return nil
}

// Match checks if the given string matches any of the filter's regex patterns.
func (f *Filter) Match(s string) bool {
	if len(f.patterns) == 0 {
		return true
	}

	for _, pattern := range f.patterns {
		if pattern.MatchString(s) {
			return true
		}
	}

	return false
}

func (f *Filter) Filter(items []string) []string {
	filtered := make([]string, 0, len(items))
	for _, item := range items {
		if f.Match(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
