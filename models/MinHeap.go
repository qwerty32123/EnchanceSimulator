package models

type MinHeap struct {
	Arr []*Node
}

func (h *MinHeap) Push(node *Node) {
	h.Arr = append(h.Arr, node)
	i := len(h.Arr) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.Arr[parent].Freq <= h.Arr[i].Freq {
			break
		}
		h.Arr[i], h.Arr[parent] = h.Arr[parent], h.Arr[i]
		i = parent
	}
}

func (h *MinHeap) Pop() *Node {
	root := h.Arr[0]
	h.Arr[0] = h.Arr[len(h.Arr)-1]
	h.Arr = h.Arr[:len(h.Arr)-1]
	i := 0
	for {
		left := 2*i + 1
		right := 2*i + 2
		min := i
		if left < len(h.Arr) && h.Arr[left].Freq < h.Arr[min].Freq {
			min = left
		}
		if right < len(h.Arr) && h.Arr[right].Freq < h.Arr[min].Freq {
			min = right
		}
		if min == i {
			break
		}
		h.Arr[i], h.Arr[min] = h.Arr[min], h.Arr[i]
		i = min
	}
	return root
}
