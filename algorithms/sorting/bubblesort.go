package sorting

/* BubbleSort
 * Sorts given slice using Bubble Sort algorithm
 * @param arr []int - unsorted slice
 * @return []int - sorted slice
 */
func BubbleSort(A []int) []int {
	for i := 1; i < len(A); i++ {
		for j := len(A) - 1; j > 0; j-- {
			if A[j] < A[j-1] {
				temp := A[j]
				A[j] = A[j-1]
				A[j-1] = temp
			}
		}
	}
	return A
}
