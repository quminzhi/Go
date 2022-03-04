package qsort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 5, 10, 7, 3, 4, 2}
	expected := []int{1, 2, 3, 4, 5, 7, 10}
	qsort(arr, 0, len(arr)-1)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Fatalf("index %d: expected to be %d, but got %d", i, expected[i], arr[i])
		}
	}
}
