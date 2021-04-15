package _136

import (
	"reflect"
	"testing"
)

func TestTableDeleteDuplicates(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 2, 1},
			expected: 1,
		},
		{
			input:    []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			input:    []int{1},
			expected: 1,
		},
	}

	for _, tests := range tests {
		output := singleNumber(tests.input)

		if !reflect.DeepEqual(output, tests.expected) {
			t.Error("Input: ", tests.input, "\n Expected: ", tests.expected, "\n Result: ", output)
		}
	}
}
