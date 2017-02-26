package main

import (
	"math/rand"
	"p_Heap"
	"fmt"
)

func main() {
	testheap := p_Heap.New()

	for i := 0; i < 20; i += 1 {
		toPush := int(rand.Uint32())
		fmt.Println(toPush)
		testheap.Push(toPush)
	}
	testheap.Visualize()

	for !testheap.Empty() {
		fmt.Println(testheap.Pop())
	}
}
