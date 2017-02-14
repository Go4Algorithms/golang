package graph


import "fmt"

type Node struct {
	Name string
	ID int
}


func newNode ( argName string) *Node {
	var r = new (Node)
	r.Name = argName
	return r
}


func describeNode( n *Node) {
	fmt.Printf("Node Name=" + n.Name)
}

