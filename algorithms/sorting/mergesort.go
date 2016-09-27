package sorting

/* MergeSort
 * Sorts given slice using Merge Sort algorithm
 * @param arr []int - unsorted slice
 * @return []int - sorted slice
 */
func MergeSort(A []int) []int {
	if len(A) > 1 {
		sliceL := A[:len(A)/2]
		sliceR := A[len(A)/2:]

		sliceL = MergeSort(sliceL)
		sliceR = MergeSort(sliceR)

		var i, j, k int
		mergedSlice := make([]int, len(sliceL)+len(sliceR))

		for i < len(sliceL) && j < len(sliceR) {
			if sliceL[i] < sliceR[j] {
				mergedSlice[k] = sliceL[i]
				i++
			} else {
				mergedSlice[k] = sliceR[j]
				j++
			}
			k++
		}

		endSlice := mergedSlice[k:]

		if i < len(sliceL) {
			copy(endSlice, sliceL[i:])
		}
		if j < len(sliceR) {
			copy(endSlice, sliceR[j:])
		}
		return mergedSlice
	}
	return A
}
