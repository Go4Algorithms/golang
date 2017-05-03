/*
Neal Wolff
Binary Search
Programed in go Language
*/

package main

import (
	"fmt"
	)

func main() {
	A := []int {0,1,3,4,5,6,7,9,11,14,15,16,17,18,20,22,25,27}
	fmt.Println("integer to find")
	var find int;
	fmt.Scan(&find)
	found,index := binarySearch(A, find, A[0], A[len(A)-1])

	if found==true {
		fmt.Println("Number", find, "was found at index",index)
	}else{
		fmt.Println("Number was not found")
	}
}

func binarySearch(A []int, value int, low int, high int) (bool, int) {
    if high < low {
        return false,0
    }
    mid := (low + high) / 2
    if A[mid] > value {
        return binarySearch(A, value, low, mid-1)
    } else if A[mid] < value {
        return binarySearch(A, value, mid+1, high)
    }
    return true, mid
}
