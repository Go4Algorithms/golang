package main

import (
	"fmt"
	"p_LinkedList"
)

func main() {
	testlist := p_LinkedList.New()

	for i := 0; i < 20; i++ {
		testlist.PushBack(i)
	}

	testlist.Visualize()
	testlist.Insert(100, 7)
	testlist.Visualize()

	for testlist.Size() > 5 {
		fmt.Println(testlist.Pull(3))
	}

	testlist.Visualize()
}
