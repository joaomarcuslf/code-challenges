package challenge04

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name       string
	input      []int
	expected01 int
	expected02 int
}

func TestExtractList(t *testing.T) {

	testCases := []TestCase{
		{
			name:       "Test with only one Repeat",
			input:      []int{1, 2, 4, 1, 3},
			expected01: 3,
			expected02: 1,
		},
		{
			name:       "Test with multiple Repeats",
			input:      []int{2, 40, 1, 3, -1, 3, 40, 5, 1},
			expected01: 3,
			expected02: 3,
		},
		{
			name:       "Test with no Repeat",
			input:      []int{2, -1, 3, 40, 5, 1},
			expected01: 6,
			expected02: 0,
		},
		{
			name:       "Test with one element",
			input:      []int{1},
			expected01: 1,
			expected02: 0,
		},
		{
			name:       "Test with no elements",
			input:      []int{},
			expected01: 0,
			expected02: 0,
		},
	}

	for _, tc := range testCases {
		u, r := ExtractList(tc.input)

		if !reflect.DeepEqual(u, tc.expected01) {
			t.Errorf("%s - (Unique) got %v, wanted %v, of type %s", tc.name, u, tc.expected01, "int")
		}

		if !reflect.DeepEqual(r, tc.expected02) {
			t.Errorf("%s - (Repeat) got %v, wanted %v, of type %s", tc.name, r, tc.expected02, "int")
		}
	}
}
