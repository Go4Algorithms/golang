package p_Heap

import "fmt"

// Heap is implemented using a slice instead of a binary tree
type Heap struct {
	arr       []int
	slice_inc int
}


// --- Public interface ---


func New() Heap {
	return Heap{make([]int, 0, 5), 5}
}

// Add a number to the heap
func (h *Heap) Push(num int) {
	if len(h.arr) == cap(h.arr) {
		h.arr = append(h.arr, make([]int, 1, h.slice_inc)...)
	} else {
		h.arr = h.arr[:len(h.arr) + 1]
	}
	h.arr[len(h.arr) - 1] = num
	h.reorderUpheap()
}

// Take a number from the heap
func (h *Heap) Pop() int {
	if len(h.arr) == 0 {
		fmt.Println("Error!  Heap.Pop() -> Heap is empty!")
		return 0
	}
	
	toReturn := h.arr[0]
	
	if len(h.arr) == 1 {
		fmt.Println("reeternin' caws sahz wuz one!")
		fmt.Println(h.arr)
		h.arr = h.arr[:0]
		return toReturn
	}
	
	h.arr[0], h.arr[len(h.arr) - 1] = swap(h.arr[0], h.arr[len(h.arr) - 1])
	h.arr = h.arr[0:len(h.arr) - 1]
	h.reorderDownheap()
	return toReturn
}

func (h *Heap) Empty() bool {
	return len(h.arr) == 0
}

func (h *Heap) Visualize() {
	fmt.Println(h.arr)
}


// --- Private helpers ---


func (h *Heap) reorderUpheap() {
	current := len(h.arr) - 1
	if current == 0 {
		return
	}
	
	// Loop until break is encountered
	for {
		parent := getParentIdx(current)
		// If we've reached the root, we're done.
		if (current <= 0) {
			break
		}
		// If the parent's value is greater than the current value, the heap is ordered and we are done
		if (h.arr[parent] >= h.arr[current]) {
			break
		}
		// Otherwise, switch parent and current and set current to parent
		h.arr[parent], h.arr[current] = swap(h.arr[parent], h.arr[current])
		current = parent
	}
}

func (h *Heap) reorderDownheap() {
	current := 0
	if current == len(h.arr) - 1 {
		return
	}
	
	// Loop until break is encountered
	for {
		// Get left child, right child, and the max between the two.
		lChild := getLChildIdx(current)
		rChild := getRChildIdx(current)
		max := lChild
		// If lChild is out of bounds for h.arr, we are currently in the last slot in the heap and it is ordered properly
		if lChild >= len(h.arr) {
			break
		// If rChild is out of bounds for h.arr but not lChild, we need to check with lChild before breaking
		} else if rChild >= len(h.arr) {
			if h.arr[lChild] > h.arr[current] {
				h.arr[lChild], h.arr[current] = swap(h.arr[lChild], h.arr[current])
			}
			break
		}
		// Find the max between left and right children
		if h.arr[rChild] > h.arr[lChild] {
			max = rChild
		}
		// If the maximum child is greater than the parent, swap the two
		if h.arr[max] > h.arr[current] {
			h.arr[max], h.arr[current] = swap(h.arr[max], h.arr[current])
			current = max
		// Otherwise, the heap is ordered and we can break from the loop
		} else {
			break
		}
	}
}

// Find the index of the parent of an index
func getParentIdx(childIdx int) int {
	return (childIdx - 1)/2
}

// Find the index of the left child of an index
func getLChildIdx(parentIdx int) int {
	return 2*parentIdx + 1
}

// Find the index of the right child of an index
func getRChildIdx(parentIdx int) int {
	return 2*parentIdx + 2
}

// Swap two numbers
func swap(a, b int) (int, int) {
	return b, a
}

