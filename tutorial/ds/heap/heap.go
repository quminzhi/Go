/*
Package heap implements min heap and common methods associated with it.
*/
package heap

import (
	"errors"
	"reflect"
)

// Heap
type Heap struct {
	size int
	vec  []int
}

/* 
InitHeap inits a heap from a vector.
*/
func (h *Heap) InitHeap(_vec []int) error {
	h.size = len(_vec)
	h.vec = _vec

	swap := reflect.Swapper(h.vec)

	// update sift-down strategy
	for i := len(_vec) - 1; i >= 0; i-- {
		root := i
		for {
			left := 2*root + 1
			right := 2*root + 2
			// case 1: no child
			if (left >= h.size) && (right >= h.size) {
				break
			}
			// case 2: only left child
			if (right >= h.size) && (h.vec[root] > h.vec[left]) {
				swap(root, left)
				root = left
				continue
			}
			// case 3: left and right child exist
			var smaller int
			if h.vec[left] < h.vec[right] {
				smaller = left
			} else {
				smaller = right
			}

			if h.vec[root] > h.vec[smaller] {
				swap(root, smaller)
				root = smaller
				continue
			}

			// otherwise
			break
		}
	}

	return nil
}

func (h Heap) Empty() (bool, error) {
	if size, err := h.Size(); err == nil {
		return size == 0, nil
	}
	return false, errors.New("fail to get size of heap")
}

func (h Heap) Size() (int, error) {
	return h.size, nil
}

func (h Heap) Top() (int, error) {
	if ret, err := h.Empty(); err == nil {
		if ret {
			return -1, errors.New("heap is empty")
		} else {
			return h.vec[0], nil
		}
	} else {
		return -1, errors.New("failed to call h.Empty()")
	}
}

func (h *Heap) Push(val int) error {
	if empty, err := h.Empty(); err == nil {
		if empty {
			// corner case: for the first element
			h.vec[0] = val
			h.size++
		} else {
			// others
			h.vec = append(h.vec, val)
			h.size++

			// update min heap
			child := h.size - 1
			parent := (child - 1) / 2
			for child != 0 && h.vec[child] < h.vec[parent] {
				h.vec[child], h.vec[parent] = h.vec[parent], h.vec[child]
				child = parent
				parent = (child - 1) / 2
			}
		}
		return nil
	} else {
		return errors.New("failed to invoke h.Empty()")
	}
}

func (h *Heap) Pop() (int, error) {
	empty, _ := h.Empty()
	if empty {
		return -1, errors.New("tried to pop from an empty heap")
	}

	// corner case: only one left in the heap
	if size, _ := h.Size(); size == 1 {
		var popped int = h.vec[0]
		h.size--
		return popped, nil
	}

	// if it is not the last one
	popped := h.vec[0]
	h.vec[0] = h.vec[h.size-1]
	h.size--
	// update: sift down
	root, left, right := 0, 0, 0
	for {
		left = 2*root + 1
		right = 2*root + 2
		if (left >= h.size) && (right >= h.size) {
			break
		}
		if (right >= h.size) && (h.vec[left] < h.vec[root]) {
			h.vec[root], h.vec[left] = h.vec[left], h.vec[root]
			root = left
			continue
		}
		var min int
		if h.vec[left] < h.vec[right] {
			min = left
		} else {
			min = right
		}
		if h.vec[root] > h.vec[min] {
			h.vec[root], h.vec[min] = h.vec[min], h.vec[root]
			root = min
			continue
		}

		// right place
		break
	}

	return popped, nil
}
