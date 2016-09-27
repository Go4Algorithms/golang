package main

import "fmt"

// Using Kadane's algorithm

// Max from math package does not support integers
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxSubArray(A []int) int {
	max := A[0]
	currentMax := A[0]
	for i := 1; i < len(A); i++ {
		max = Max(A[i], max+A[i])
		currentMax = Max(currentMax, max)
	}
	return currentMax
}
func main() {
	var c = []int{-1, -2, -3, -4, -5}
	fmt.Println(MaxSubArray(c))
}
