package xregexp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilter_AddPatternAndMatch(t *testing.T) {
	f := &Filter{}
	require.NoError(t, f.AddPattern("foo"), "AddPattern should not fail")
	assert.True(t, f.Match("foobar"), "Expected Match to return true for 'foobar'")
	assert.False(t, f.Match("bar"), "Expected Match to return false for 'bar'")
}

func TestFilter_AddPatterns(t *testing.T) {
	f := &Filter{}
	patterns := []string{"foo", "bar"}
	require.NoError(t, f.AddPatterns(patterns...), "AddPatterns should not fail")
	assert.True(t, f.Match("foo"), "Expected Match to return true for 'foo'")
	assert.True(t, f.Match("bar"), "Expected Match to return true for 'bar'")
	assert.False(t, f.Match("baz"), "Expected Match to return false for 'baz'")
}

func ExampleFilter_Match() {
	f := &Filter{}
	_ = f.AddPattern("foo")
	_ = f.AddPattern("bar")
	fmt.Println(f.Match("foo")) // true
	fmt.Println(f.Match("baz")) // false
	// Output:
	// true
	// false
}
