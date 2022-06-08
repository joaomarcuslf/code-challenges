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

// Time: 0.884s
func TestSortArray(t *testing.T) {
	got := SortArrayMergeSort(Input())

	if !reflect.DeepEqual(got, Output()) {
		t.Errorf("got %v, wanted %v, of type %s", got, Output(), "int")
	}
}

// Time: 0.320s
func TestSortArrayCountingSort(t *testing.T) {
	got := SortArrayCountingSort(Input())

	if !reflect.DeepEqual(got, Output()) {
		t.Errorf("got %v, wanted %v, of type %s", got, Output(), "int")
	}
}

// Time: 0.315s
func TestSortArrayQuickSort(t *testing.T) {
	got := SortArrayQuickSort(Input())

	if !reflect.DeepEqual(got, Output()) {
		t.Errorf("got %v, wanted %v, of type %s", got, Output(), "int")
	}
}

func TestSortArrayBubbleSort(t *testing.T) {
	got := SortArrayBubbleSort(Input())

	if !reflect.DeepEqual(got, Output()) {
		t.Errorf("got %v, wanted %v, of type %s", got, Output(), "int")
	}
}
