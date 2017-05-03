package main

import (
	"fmt"
)
func main() {
	var root *Node
	A := []int {0,17,2,16,7,18,19,6,3,4,1,20,5,14,12,15,10,9,13,8,11}
	for _, data := range A {
		root = insert(root, data)
	}

	inOrder(root,true)

}


type Node struct {
	Data int //node data
	Left, Right *Node //left and right child nodes
}

//insert data into the tree
func insert(root *Node, data int) *Node{
	//if the tree has not been created yet, insert in to root.
	if root == nil {
		root = &Node{data, nil, nil}
		return root
	} 

	//if the data is less than the data contained in the root
	//recursivly insert into the left node.
	if data < root.Data {
		root.Left = insert(root.Left, data)
	} 

	//if the data is greater than the data contained in the root,
	//recursivly insert into the right node
	if data > root.Data {
		root.Right = insert(root.Right, data)
	}
	return root
}

//in order tree traversal
func inOrder(root *Node, LoR bool){
	if root == nil {
		return
	}
	if LoR {
		fmt.Println(root.Data, " (left)")
	}else {
		fmt.Println(root.Data, " (right)")
	}

	inOrder(root.Left, true)
	inOrder(root.Right, false)
}
