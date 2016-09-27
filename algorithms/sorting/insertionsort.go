package sorting

/* InsertionSort
 * Sorts given slice using Insertion Sort algorithm
 * @param arr []int - unsorted slice
 * @return []int - sorted slice
 */
func InsertionSort(A []int) []int {

	for j := 1; j < len(A); j++ {
		key := A[j]

		i := j - 1
		for ; i >= 0 && A[i] > key; i-- {
			A[i+1] = A[i]
		}
		A[i+1] = key
	}
	return A
}
