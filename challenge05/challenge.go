package challenge05

import (
	"sort"
)

type BinarySearchTree struct {
	value int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func _fromArrayToBinarySearchTree(input []int) *BinarySearchTree {
	if len(input) == 0 {
		return nil
	}

	mid := len(input) / 2
	root := &BinarySearchTree{input[mid], nil, nil}

	root.left = _fromArrayToBinarySearchTree(input[:mid])
	root.right = _fromArrayToBinarySearchTree(input[mid+1:])

	return root
}

func FromArrayToBinarySearchTree(input []int) *BinarySearchTree {
	sort.Ints(input)

	return _fromArrayToBinarySearchTree(input)
}

func SearchBinarySearchTree(root *BinarySearchTree, value int) bool {
	if root == nil {
		return false
	}

	if root.value == value {
		return true
	}

	if value < root.value {
		return SearchBinarySearchTree(root.left, value)
	}

	return SearchBinarySearchTree(root.right, value)
}
