package challenge01

import (
	"reflect"
	"testing"
)

func Input() []int {
	return []int{0, 2, 2, 1, 1, 0, 1, 0, 0, 2, 1, 1, 0, 1, 2, 1, 1, 1, 2, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 2, 1, 0, 0, 1, 0, 2, 1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 2, 1, 2, 0, 0, 1, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 2, 0, 2, 1, 1, 2, 0, 1, 0, 0, 1, 1, 1}
}

func Output() []int {
	return []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
}

type TestCase struct {
	name     string
	input    []int
	expected []int
}

var testCases []TestCase = []TestCase{
	{
		name:     "Testing small array",
		input:    []int{1, 0, 2},
		expected: []int{0, 1, 2},
	},
	{
		name:     "Testing big array",
		input:    []int{0, 2, 2, 1, 1, 0, 1, 0, 0, 2, 1, 1, 0, 1, 2, 1, 1, 1, 2, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 2, 1, 0, 0, 1, 0, 2, 1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 2, 1, 2, 0, 0, 1, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 2, 0, 2, 1, 1, 2, 0, 1, 0, 0, 1, 1, 1},
		expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},
	{
		name:     "Testing already sorted medium array",
		input:    []int{0, 0, 1, 2, 2},
		expected: []int{0, 0, 1, 2, 2},
	},
	{
		name:     "Testing reverse sorted medium array",
		input:    []int{2, 2, 1, 0, 0},
		expected: []int{0, 0, 1, 2, 2},
	},
	{
		name:     "Testing with only one element",
		input:    []int{0},
		expected: []int{0},
	},
	{
		name:     "Testing with only no element",
		input:    []int{},
		expected: []int{},
	},
}

func TestSortArrayMergeSort(t *testing.T) {

	for _, tc := range testCases {
		output := SortArrayMergeSort(tc.input)

		if (len(output) > 0 && len(tc.expected) > 0) && !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "[]int")
		}
	}
}

func TestSortArrayCountingSort(t *testing.T) {

	for _, tc := range testCases {
		output := SortArrayCountingSort(tc.input)

		if (len(output) > 0 && len(tc.expected) > 0) && !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "[]int")
		}
	}
}

func TestSortArrayQuickSort(t *testing.T) {

	for _, tc := range testCases {
		output := SortArrayQuickSort(tc.input)

		if (len(output) > 0 && len(tc.expected) > 0) && !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "[]int")
		}
	}
}

func TestSortArrayBubbleSort(t *testing.T) {

	for _, tc := range testCases {
		output := SortArrayBubbleSort(tc.input)

		if (len(output) > 0 && len(tc.expected) > 0) && !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "[]int")
		}
	}
}
