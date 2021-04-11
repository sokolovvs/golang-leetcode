package _66

import (
	"reflect"
	"testing"
)

func TestTableDeleteDuplicates(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			input:    []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			input:    []int{0},
			expected: []int{1},
		},
		{
			input:    []int{9},
			expected: []int{1, 0},
		},
	}

	for _, tests := range tests {
		output := plusOne(tests.input)

		if !reflect.DeepEqual(output, tests.expected) {
			t.Error("Input: ", tests.input, "\n Expected: ", tests.expected, "\n Result: ", output)
		}
	}
}
