/*
Neal Wolff
Insertion Sort
Programed in go Language
*/

package main

import (
	"fmt"
	)

func main() {
	A := []int {17,2,16,7,18,19,6,3,4,1,20,5,14,12,15,10,9,13,8,11}
	fmt.Println(A)
	heapsort(A)
	fmt.Println(A)
}

func downheap (A []int, n int, i int) {
    for true {
        j := max(A, n, i, 2*i+1, 2*i+2);
        if (j == i) {
            break;
        }
        t := A[i]
        A[i] = A[j]
        A[j] = t
        i = j
    }
}
 
func heapsort (A []int){
	n := len(A);
	i:=0;
	for i =(n-2)/2; i >= 0; i--{
		downheap(A,n,i)
	}
	for i = 0; i < n; i++{
		t := A[n-i-1]
		A[n-i-1] = A[0]
		A[0] = t
		downheap(A, n-i-1, 0)
	}
}

func max (A []int,n int,i int,j int,k int) int{
	m := i
	if (j < n && A[j] > A[m]){
		m = j
	}
	if (k < n && A[k] > A[m]) {
		m = k
	}
	return m;
}
