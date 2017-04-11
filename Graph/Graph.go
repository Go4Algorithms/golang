package graph

import (
	"fmt"
)


type Graph struct {
	Name string
	Desc string
	Nodes [] Node
	Edges [] Edge
	Weighted bool
	Directed bool
	Matrix [][] int
}

func NewGraph(name string, directed bool, weighted bool, nodes[]Node,edges[]Edge)*Graph {
	var g = new (Graph)
	g.Name = name
	g.Directed = directed
	g.Nodes = nodes
	g.Edges = edges
	g.Weighted = weighted
	if weighted {
		g.Matrix = createWeightedGraphMatrix(g)
	}
	return g
}

//creates a matrix with all the weights. Useful for shortest path and MST
func createWeightedGraphMatrix (g *Graph)[][]int {
    gm := make([][]int,len(g.Nodes))
    for i:= range gm {
        gm[i] = make([]int,len(g.Nodes))
    }
    for row:= range gm {
        for col:= range gm {
            gm[row][col] = 0
        }
    }
    for _,e:= range g.Edges {
	    if !g.Directed{
		gm[e.To.ID][e.From.ID] = e.Weight 
	}
        gm[e.From.ID][e.To.ID] = e.Weight
    }
    return gm
}

func PrintNodes ( nodes [] Node) {
	fmt.Println("Nodes in the graph:")
	for _,n := range nodes {
		fmt.Printf(" %s,", n.Name)
	}
	fmt.Println("")
}

func PrintEdges ( edges [] Edge) {
	fmt.Println("Edges in the graph:")
	for _,e := range edges {
		fmt.Printf(" %s,", e.Name)
	}
	fmt.Println("")
}

func FindEdge ( eArg Edge, g * Graph ) ( bool ) {
	var ret = false
	for  _,e := range g.Edges {
		if (  e == eArg ) {
			return  true
		}
	}
	return ret
}

func DescribeGraph( g *Graph) {

	fmt.Println("Graph name =", g.Name)
	PrintNodes( g.Nodes )
	fmt.Println("")
	PrintEdges ( g.Edges )
}


func DescribeMatrix ( g *Graph ) {

	fmt.Println("Graph name =", g.Name)
	fmt.Println("Nodes in the graph:")
	for _,n := range g.Nodes {
		fmt.Printf(" %s,", n.Name)
	}
	fmt.Println("")

	fmt.Println("The Matrix Nodes in the graph is:")

	tm := g.Matrix
	for i := range tm {
		for j := range tm[i] {
			fmt.Printf( "  m[%d][%d] = %d ", i,j,tm[i][j])
		}
		fmt.Println("")
	}

}
