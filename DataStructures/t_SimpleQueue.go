package main

import (
	"fmt"
	"p_SimpleQueue"
)

func main() {
	testqueue := p_SimpleQueue.New()
	
	for i := 0; i < 100; i++ {
		testqueue.Push(i)
	}
	
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(" ", testqueue.Pop())
		}
		fmt.Print("\n")
	}
}
