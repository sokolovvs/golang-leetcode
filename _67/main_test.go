package _67

import (
	"reflect"
	"testing"
)

func TestTableDeleteDuplicates(t *testing.T) {
	var tests = []struct {
		addendum      string
		otherAddendum string
		expected      string
	}{
		{
			addendum:      "11",
			otherAddendum: "1",
			expected:      "100",
		},
		{
			addendum:      "1010",
			otherAddendum: "1011",
			expected:      "10101",
		},
		{
			addendum:      "1111",
			otherAddendum: "1111",
			expected:      "11110",
		},
		{
			addendum:      "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			otherAddendum: "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			expected:      "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
		{
			addendum:      "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			otherAddendum: "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			expected:      "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for _, tests := range tests {
		output := addBinary(tests.addendum, tests.otherAddendum)

		if !reflect.DeepEqual(output, tests.expected) {
			t.Error("Addendums: ", tests.addendum, tests.otherAddendum, "\n Expected: ", tests.expected, "\n Result: ", output)
		}
	}
}
