package main

import (
	"fmt"

	"minzhi.io/heap"
)

func read(n int) ([]int, error) {
	in := make([]int, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func main() {
	fmt.Println("please input the size of vec: (ex> 3)")
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("error in reading the size of vec")
	} else {
		fmt.Println("okay and then give the value of each element: (ex> 2 3 1)")
		vec, _ := read(n)
		var h heap.Heap
		h.InitHeap(vec)
		for empty, _ := h.Empty(); !empty; empty, _ = h.Empty() {
			top, _ := h.Top()
			fmt.Println(top)
			h.Pop()
		}
	}
}
