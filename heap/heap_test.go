package heap_test

import (
	"testing"

	"github.com/victorprb/data-structures/heap"
)

func TestHeap_Add(t *testing.T) {
	h := heap.MaxHeap{}

	values := []int{10, 20, 30, 40, 5}

	for _, v := range values {
		h.Add(v)
	}

	want := []int{40, 30, 20, 10, 5}
	for _, v := range want {
		got, err := h.Get()
		if err != nil {
			t.Errorf("should not return error")
		}

		if got != v {
			t.Errorf("got %d, want %d", got, v)
		}
	}
}
