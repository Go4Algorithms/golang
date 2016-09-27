package sorting

/* SelectionSort
 * Sorts given slice using Selection Sort algorithm
 * @param arr []int - unsorted slice
 * @return []int - sorted slice
 */
func SelectionSort(A []int) []int {
	for i := 0; i < len(A); i++ {
		min := i

		for j := i + 1; j < len(A); j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		temp := A[i]
		A[i] = A[min]
		A[min] = temp
	}
	return A
}
