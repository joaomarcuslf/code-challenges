package main

import (
	"encoding/json"
	"fmt"
)

type BinTree struct {
	Value int      `json:"value"`
	Left  *BinTree `json:"left,omitempty"`
	Right *BinTree `json:"right,omitempty"`
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *BinTree) Insert(value int) {
	if value < t.Value {
		if t.Left == nil {
			t.Left = &BinTree{Value: value}
		} else {
			t.Left.Insert(value)
		}
		return
	} else {
		if t.Right == nil {
			t.Right = &BinTree{Value: value}
		} else {
			t.Right.Insert(value)
		}
		return
	}
}

func (t *BinTree) Search(needle int) bool {
	if t.Value == needle {
		return true
	}
	if needle < t.Value && t.Left != nil {
		return t.Left.Search(needle)
	}
	if needle > t.Value && t.Right != nil {
		return t.Right.Search(needle)
	}
	return false
}

func (t *BinTree) Height() int {
	if t == nil {
		return 0
	}
	return 1 + max(t.Left.Height(), t.Right.Height())
}

func (t *BinTree) Balance() int {
	if t == nil {
		return 0
	}
	return t.Left.Height() - t.Right.Height()
}

func NewBinaryTree(K []int) *BinTree {
	tree := &BinTree{
		Value: K[0],
	}

	for i := range K {
		if i > 0 {
			fmt.Println(K[i])

			tree.Insert(K[i])
		}
	}
	return tree
}

func SearchBinarySearchTree(tree *BinTree, needle int) bool {
	return tree.Search(needle)
}

func main() {
	K := []int{0, 2, 3, 1, 4, 6, 7, 7}
	tree := NewBinaryTree(K)
	treeJSON, _ := json.Marshal(tree)
	fmt.Println(string(treeJSON))
	fmt.Println(SearchBinarySearchTree(tree, 7))
}
