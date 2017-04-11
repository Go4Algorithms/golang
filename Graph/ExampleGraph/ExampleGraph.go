package main

import (
	"fmt"
	"github.com/Go4Algorithms/golang/Graph"
)

//this sample graph can be found at www.siddarthareddy.weebly.com/uploads/2/8/7/9/28799429/4999443.png?468
func main() {
	nodes := []graph.Node{*graph.NewNode("A",0),*graph.NewNode("B",1),*graph.NewNode("C",2),*graph.NewNode("D",3),*graph.NewNode("S",4)}
	edges := []graph.Edge{*graph.NewWeightedEdge(&nodes[0],&nodes[1],"A->B",1,0),*graph.NewWeightedEdge(&nodes[0],&nodes[2],"A->C",2,1),*graph.NewWeightedEdge(&nodes[1],&nodes[3],"B->D",4,2),
		*graph.NewWeightedEdge(&nodes[2],&nodes[0],"C->A",3,3),*graph.NewWeightedEdge(&nodes[2],&nodes[1],"C->B",9,4),*graph.NewWeightedEdge(&nodes[2],&nodes[3],"C->D",2,5),
		*graph.NewWeightedEdge(&nodes[3],&nodes[1],"D->B",6,6),*graph.NewWeightedEdge(&nodes[3],&nodes[4],"D->S",7,7),*graph.NewWeightedEdge(&nodes[4],&nodes[0],"S->A",10,8),
		*graph.NewWeightedEdge(&nodes[4],&nodes[2],"S->C",5,9)}
	g := graph.NewGraph("Sample",true,true,nodes,edges)
	f := graph.FloydWarshall(g)
	fmt.Println("Floyd Warshall")
	for i:=0; i<len(f); i++ {
		for j:=0; j<len(f[0]); j++ {
			fmt.Print(f[i][j]," ")
		}
		fmt.Println()
	}
	b := graph.BreadthFirstSearch(g,0)
	fmt.Println("Breadth First Search")
	for i:=0; i<len(b); i++ {
		fmt.Println(b[i])
	}
	d := graph.DepthFirstSearch(g,0)
	fmt.Println("Depth First Search")
	for i:=0; i<len(d); i++ {
		fmt.Println(d[i])
	}
	dj := graph.DijkstraDistance(g,0)
	fmt.Println("Dijkstra Distances")
	for i:=0; i<len(dj); i++{
		fmt.Println(dj[i])
	}
	dk := graph.DijkstraEdges(g,&g.Nodes[0])
	fmt.Println("Dijkstra Edges")
	for i:=0; i<len(dk); i++{
		fmt.Println(dk[i].Name)
	}
	g.Encode("example")
}

