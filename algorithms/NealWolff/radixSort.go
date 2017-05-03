package main

import (
	"fmt"
)

func main() {
	A := []int {17,2,16,7,18,19,6,3,4,1,20,5,14,12,15,10,9,13,8,11}
	fmt.Println(A)
	radixSort(A)
	fmt.Println(A)
}
// Radix Sort
func radixSort(A []int) {
	bSize :=10
	maxNum := max(A)
	n := len(A)
	bsD := 1
	semiSorted := make([]int, n, n)


	for maxNum/bsD > 0 {
		bucket := make([]int,bSize)

		for i := 0; i < n; i++ {
			bucket[(A[i] / bsD) % bSize]++
		}

		for i := 1; i < bSize; i++ {
			bucket[i] += bucket[i-1] 
		}

		for i := n - 1; i >= 0; i-- {
			bucket[(A[i] / bsD) % bSize]--
			semiSorted[bucket[(A[i] / bsD) % bSize]] = A[i]
		}

		for i := 0; i < n; i++ {
			A[i] = semiSorted[i]
		}
  
		bsD *= bSize 
	}
}


func max(A []int) int {
	maxNum := 0
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] > maxNum {
			maxNum = A[i]
		}
	}
	return maxNum
}

