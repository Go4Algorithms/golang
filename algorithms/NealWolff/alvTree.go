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
	Data int
	Height int
	Left, Right *Node
}

func leftRotate(root *Node) *Node {
	node := root.Right
	root.Right = node.Left
	node.Left = root

	root.Height = max(height(root.Left), height(root.Right)) + 1
	node.Height = max(height(node.Right), height(node.Left)) + 1
	return node
}

func leftRightRotate(root *Node) *Node {
	root.Left = leftRotate(root.Left)
	root = rightRotate(root)
	return  root
}

func rightRotate(root *Node) *Node {
	node := root.Left
	root.Left = node.Right
	node.Right = root
	root.Height = max(height(root.Left), height(root.Right)) + 1
	node.Height = max(height(node.Left), height(node.Right)) + 1
	return node
}

func rightLeftRotate(root *Node) *Node {
	root.Right = rightRotate(root.Right)
	root = leftRotate(root)
	return  root
}

func height(root *Node) int {
	if root != nil {
		return root.Height
	}
	return -1
}

//insert data into the tree
func insert(root *Node, data int) *Node {
	if root == nil {
		root = &Node{data, 0, nil, nil}
		root.Height = max(height(root.Left), height(root.Right)) + 1
		return root
	} 

	if data < root.Data {
		root.Left = insert(root.Left, data)
		if height(root.Left)-height(root.Right) == 2 {
			if data < root.Left.Data {
				root = rightRotate(root)
			} else {
				root = leftRightRotate(root)
			}
		}
	} 

	if data > root.Data {
		root.Right = insert(root.Right, data)
		if height(root.Right)-height(root.Left) == 2 {
			if data > root.Right.Data {
				root = leftRotate(root)
			} else {
				root = rightLeftRotate(root)
			}
		}
	}

	root.Height = max(height(root.Left), height(root.Right)) + 1
	return root
}

//find the larger number
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


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
