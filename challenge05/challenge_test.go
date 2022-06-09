package challenge05

import (
	"reflect"
	"testing"
)

func TestFromArrayToBinarySearchTree(t *testing.T) {
	type TestCase struct {
		name     string
		input    []int
		expected *BinarySearchTree
	}

	testCases := []TestCase{
		{
			name:     "Test empty array",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "Test small Array",
			input:    []int{2, 3, 1},
			expected: &BinarySearchTree{2, &BinarySearchTree{1, nil, nil}, &BinarySearchTree{3, nil, nil}},
		},
		{
			name:     "Test small sorted Array",
			input:    []int{1, 2, 3},
			expected: &BinarySearchTree{2, &BinarySearchTree{1, nil, nil}, &BinarySearchTree{3, nil, nil}},
		},
		{
			name:     "Test small reverse sorted Array",
			input:    []int{3, 2, 1},
			expected: &BinarySearchTree{2, &BinarySearchTree{1, nil, nil}, &BinarySearchTree{3, nil, nil}},
		},
	}

	for _, tc := range testCases {
		output := FromArrayToBinarySearchTree(tc.input)

		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "*BinarySearchTree")
		}
	}
}

func TestSearchBinarySearchTree(t *testing.T) {
	type TestCase struct {
		name     string
		input    int
		expected bool
	}

	bigTree := FromArrayToBinarySearchTree([]int{1, 5, 12, 2, 67, 31, 56, 3, 78, 4, 5, 7, 6, 9, 10, 93, 67, 54, 37, 86})

	testCases := []TestCase{
		{
			name:     "Test existing middle element",
			input:    56,
			expected: true,
		},
		{
			name:     "Test existing last element",
			input:    93,
			expected: true,
		},
		{
			name:     "Test non-existing lesser element",
			input:    0,
			expected: false,
		},
		{
			name:     "Test non-existing big element",
			input:    190,
			expected: false,
		},
	}

	for _, tc := range testCases {
		output := SearchBinarySearchTree(bigTree, tc.input)

		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "*BinarySearchTree")
		}
	}
}
