package heap

import "errors"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Add(v int) {
	h.array = append(h.array, v)

	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Get() (int, error) {
	if len(h.array) == 0 {
		return 0, errors.New("empty heap")
	}

	head := h.array[0]

	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.heapifyDown(0)

	return head, nil
}

func (h *MaxHeap) swap(a, b int) {
	h.array[a], h.array[b] = h.array[b], h.array[a]
}

func (h *MaxHeap) heapifyUp(child int) {
	for h.array[parent(child)] < h.array[child] {
		h.swap(parent(child), child)

		child = parent(child)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	last := len(h.array) - 1
	left, right := leftChild(index), rightChild(index)
	compare := 0

	for left <= last {
		if left == last || h.array[left] > h.array[right] {
			compare = left
		} else {
			compare = right
		}

		if h.array[index] < h.array[compare] {
			h.swap(index, compare)
			index = compare
			left, right = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}
