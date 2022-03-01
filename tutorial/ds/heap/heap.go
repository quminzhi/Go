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
	if ret, err := h.Size(); err == nil {
		return ret == 0, nil
	}
	return false, errors.New("fail to get size of heap")
}

func (h Heap) Size() (int, error) {
	return len(h.vec), nil
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

func Pop() (int, error) {
	return 0, nil
}
