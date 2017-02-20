package p_LinkedList

import "fmt"

type LinkedList struct {
	first *node
	last  *node
	size  int
}

type node struct {
	next     *node
	previous *node
	val      int
}

// Constructer b/c I like them.  Not really necessary here, calling p_LinkedList.LinkedList{} will accomplish the same thing.
func New() LinkedList {
	newLL := LinkedList{nil, nil, 0}
	return newLL
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) PushBack(num int) {
	newNode := node{nil, ll.last, num}
	if ll.first == nil {
		ll.first = &newNode
	} else {
		ll.last.next = &newNode
	}
	ll.last = &newNode
	ll.size++
}

func (ll *LinkedList) PushFront(num int) {
	newNode := node{ll.first, nil, num}
	if ll.last == nil {
		ll.first = &newNode
	} else {
		ll.first.previous = &newNode
	}
	ll.first = &newNode
	ll.size++
}

func (ll *LinkedList) PopBack() int {
	toPop := ll.last
	if ll.size > 1 {
		ll.last = toPop.previous
		ll.last.next = nil
	} else {
		ll.last = nil
	}
	ll.size--
	return toPop.val
}

func (ll *LinkedList) PopFront() int {
	toPop := ll.first
	if ll.size > 1 {
		ll.first = toPop.next
		ll.first.previous = nil
	} else {
		ll.first = nil
	}
	ll.size--
	return toPop.val
}

func (ll *LinkedList) Insert(num int, pos int) {
	if (pos < 0) || (pos > ll.size) {
		fmt.Println("Error! Index passed to Linkedlist.Insert must be on the range [0, LinkedList.size]! Current size : ", ll.size, " - Passed index : ", pos)
		return
	}
	if pos == ll.size {
		ll.PushBack(num)
		return
	} else if pos == 0 {
		ll.PushFront(num)
		return
	}
	atPosCurrently := ll.traverse(pos)
	newNode := node{}
	newNode.next = atPosCurrently
	newNode.previous = atPosCurrently.previous
	atPosCurrently.previous.next = &newNode
	atPosCurrently.previous = &newNode
	newNode.val = num
	ll.size++
}

func (ll *LinkedList) Pull(pos int) int {
	// Once the references to 'toPull' are gone, garbage collection will delete 'toDelete'
	if (pos < 0) || (pos > ll.size-1) {
		fmt.Println("Error! Index passed to LinkedList.Pull must be on the range [0, LinkedList.size)! Current size : ", ll.size, " - Passed index : ", pos)
		return 0
	}
	toPull := ll.traverse(pos)
	if pos == 0 {
		return ll.PopFront()
	} else if pos == ll.size-1 {
		return ll.PopBack()
	}
	toPull.previous.next = toPull.next
	toPull.next.previous = toPull.previous
	ll.size--
	return toPull.val
}

func (ll *LinkedList) Delete(pos int) {
	// Once the references to 'toDelete' are gone, garbage collection will delete 'toDelete'
	if (pos < 0) || (pos > ll.size-1) {
		fmt.Println("Error! Index passed to LinkedList.Delete must be on the range [0, LinkedList.size)! Current size : ", ll.size, " - Passed index : ", pos)
		return
	}
	toDelete := ll.traverse(pos)
	if pos > 0 {
		toDelete.previous.next = toDelete.next
	} else {
		ll.first = toDelete.next
	}
	if pos < ll.size-1 {
		toDelete.next.previous = toDelete.previous
	} else {
		ll.last = toDelete.previous
	}
	ll.size--
}

func (ll *LinkedList) traverse(to int) *node {
	if to < ll.size {
		curNode := ll.first
		for a := 0; a < to; a++ {
			curNode = curNode.next
		}
		return curNode
	} else {
		fmt.Println("Error! Index passed to LinkedList.traverse must be less than LinkedList.size!")
	}
	return nil
}

// DEBUG! //

func (ll *LinkedList) Visualize() {
	curNode := ll.first
	i := 0
	for {
		fmt.Println(i, curNode)
		if curNode.next == nil {
			break
		}
		i++
		curNode = curNode.next
	}
}
