package main

import (
	"fmt"
	"p_SimpleStack"
)

func main() {
	teststack := p_SimpleStack.New()
	
	for i := 0; i < 100; i++ {
		teststack.Push(i)
	}
	
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(" ", teststack.Pop())
		}
		fmt.Print("\n")
	}
}
