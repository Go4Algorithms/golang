package main

import "fmt"

func InsertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
	return A
}

func main() {
    A  := []int{9, 0, 2, 1, 0}
    fmt.Println(InsertionSort(A))
}
