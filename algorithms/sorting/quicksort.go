package sorting

import "math/rand"

/* QuickSort
 * Sorts given slice using Quick Sort algorithm
 * @param arr []int - unsorted slice
 * @return []int - sorted slice
 */
func QuickSort(A []int) []int {
	if len(A) < 2 {
		return A
	}

	left, right := 0, len(A)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(A)

	// Move the pivot to the right
	A[pivotIndex], A[right] = A[right], A[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range A {
		if A[i] < A[right] {
			A[i], A[left] = A[left], A[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	A[left], A[right] = A[right], A[left]

	// Go down the rabbit hole
	QuickSort(A[:left])
	QuickSort(A[left+1:])

	return A
}
