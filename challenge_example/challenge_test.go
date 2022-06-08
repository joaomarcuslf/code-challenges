package challenge_example

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name     string
	input    []int
	expected int
}

func TestMethod(t *testing.T) {

	testCases := []TestCase{
		{
			name:     "TestCase1",
			input:    []int{},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		output := Method(tc.input)

		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "int")
		}
	}
}
