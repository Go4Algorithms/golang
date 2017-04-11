package graph


import "fmt"

type Node struct {
	Name string
	ID int
}


func NewNode ( argName string, id int) *Node {
	var r = new (Node)
	r.Name = argName
	r.ID = id
	return r
}


func DescribeNode( n *Node) {
	fmt.Printf("Node Name=" + n.Name)
}

