package qsort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 5, 10, 7, 3, 4, 2}
	expected := []int{1, 2, 3, 4, 5, 7, 10}
	qsort(arr, 0, len(arr)-1)

	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Fatalf("index %d: expected to be %d, but got %d", i, expected[i], arr[i])
		}
	}

	arr2 := []int{5, 1, 1, 2, 0, 0}
	expected2 := []int{0, 0, 1, 1, 2, 5}
	qsort(arr2, 0, len(arr2)-1)

	for i := 0; i < len(arr2); i++ {
		if arr2[i] != expected2[i] {
			t.Fatalf("index %d: expected to be %d, but got %d", i, expected2[i], arr2[i])
		}
	}

	var size int = 50000000
	arr3 := make([]int, size)
	expected3 := make([]int, size)
	for i := 0; i < size; i++ {
		arr3[i], expected3[i] = i, i
	}
	qsort(arr3, 0, len(arr3)-1)
	for i := 0; i < len(arr2); i++ {
		if arr3[i] != expected3[i] {
			t.Fatalf("index %d: expected to be %d, but got %d", i, expected2[i], arr2[i])
		}
	}
}
