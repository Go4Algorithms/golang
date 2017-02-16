package graph
//methods to add nodes and edges are needed


// go get github.com/gonum/matrix
// userID = JuanVargas
// pwd = CAEHouse+Number
import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	//  "github.com/gonum/matrix/mat64"
)

/*Graph is a type (class) that contains two sets and one attribute.
The sets are {Nodes} and {Edgees} implemented as arrays.
The attribute for now is just the Type, which could be Directed or Undirected
*/
type Graph struct {
	Name string
	Desc string
	Nodes [] Node
	Edges [] Edge
	Directed bool
	Matrix [][] float64
}


func CreateVisitToAsia() ( *Graph ) {
	var g = new (Graph)
	g.Name = "Visit to Asia"

	A := new (Node) ; A.Name = "A"

	S := newNode("S") ; 	T := newNode("T")
	C := newNode("C") ; 	B := newNode("B")
	O := newNode("O") ; 	X := newNode("X")
	D := newNode("D")

	// Interesting way to init g.Nodes
	g.Nodes = [] Node {*A,*S,*T,*C,*B,*O,*X,*D}

	// find a way to initlaize edges differently,maybe with append

	AT := newEdge(A,T,"A->T") ; 	SC := newEdge(S,C,"S->C")
	SB := newEdge(S,B,"S->B") ; 	TO := newEdge(T,O,"T->0")
	CO := newEdge(C,O,"C->O") ; 	OX := newEdge(O,X,"O->X")
	OD := newEdge(O,D,"O->D") ; 	BD := newEdge(B,D,"B->D")

	g.Edges = [] Edge {*AT,*SC,*SB,*TO,*CO,*OX,*OD,*BD}

	// 	append( g.Edges, *SC)

	return g
}


func Make_N_Nodes ( nNodes int ) ([] Node) {
	nodes := make ( []Node, nNodes)
	for i := range nodes {
		nodes[i] = * newNode("N" + strconv.Itoa(i))
	}
	return nodes
}

func Make_N_Edges ( nEdges int ) ([] Edge) {
	edges := make ( []Edge, nEdges)
	for i := range edges {
		var e = * new (Edge)
		e.Name = "E"+strconv.Itoa(i)
		edges[i] = e
	}
	return edges
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



func CreateRandomGraph ( nNodes, nEdges int)  ( *Graph) {
	var nodes = Make_N_Nodes( nNodes)
	var edges = Make_N_Edges( nEdges )
	var g = new (Graph)
	g.Nodes = nodes
	g.Edges = edges
	var s = fmt.Sprintf("Random Graph with %d Nodes and %d Edges",nNodes,nEdges)
	g.Name = s
	return g
}


func CreateRandomConnections ( min, max int, g * Graph) {
	// PrintNodes(g.Nodes)
	// PrintEdges(g.Edges)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 20; i++  {
		fmt.Println( r1.Intn(10),",")

	}
}


func CreateGraphfromFile ( aFileName string ) *Graph {
	var g = new (Graph)
	g.Name = "graph from file " + aFileName

	return g
}


func CreatePageRankSample() ( *Graph) {
	var g = new (Graph)
	g.Name = "Page Rank Example, with 6 pages"

	P1 := newNode("P1") ; P2 := newNode("P2") ; P3 := newNode("P3") ;
	P4 := newNode("P4") ; P5 := newNode("P5") ; P6 := newNode("P6") ;
	g.Nodes = [] Node { *P1, *P2, *P3, *P4, *P5, *P6 }

	P1P2 := newEdge(P1,P2,"P1->P2") ; P1P3 := newEdge(P1,P3, "P1->P3")
	P3P1 := newEdge(P3,P1,"P3->P1") ; P3P2 := newEdge(P3,P2, "P3->P2") ; P3P5 := newEdge(P3,P5 ,"P3->P5")
	P4P5 := newEdge(P4,P5,"P4->P5") ; P4P6 := newEdge(P4,P6, "P4->P6")
	P5P4 := newEdge(P5,P4,"P5->P4") ; P5P6 := newEdge(P5,P6, "P5->P6")
	P6P4 := newEdge(P6,P4,"P6->P4")

	g.Edges = [] Edge {*P1P2, *P1P3, *P3P1, *P3P2, *P3P5, *P4P5, *P4P6, *P5P4, *P5P6, *P6P4 }

	return g
}


//creates a matrix with all the weights. Useful for shortest path and MST
func CreateWeightedGraphMatrix (g *Graph)[][]int {
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
        gm[e.From.ID][e.To.ID] = e.Weight
    }
    return gm
}

func CreateTransitionMatrix ( g *Graph ) {
	var rows = len ( g.Nodes )
	var cols = rows

	// m := mat64.NewDense( rows, cols,nil)

	tm := make ( [][] float64, rows)
	for i := range tm {
		tm[i] = make ( [] float64, cols)
	}

	g.Matrix = tm
}


func CreateTransitionMatrix_t ( g *Graph ) {
	var rows = len ( g.Nodes )
	var cols = rows

	// m := mat64.NewDense( rows, cols,nil)


	var mm = [][] float64{
		{0.0, 0.5, 0.5, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		{0.3, 0.3, 0.0, 0.0, 0.3, 0.0},
		{0.0, 0.0, 0.0, 0.0, 0.5, 0.5},
		{0.0, 0.0, 0.0, 0.5, 0.0, 0.5},
		{0.0, 0.0, 0.0, 1.0, 0.0, 0.0}}

	//var row1 = [] float64 { 0.0, 0.5, 0.5, 0.0, 0.0, 0.0}
	//var row2 = [] float64 { 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	//var row3 = [] float64 { 0.3, 0.3, 0.0, 0.0, 0.3, 0.0}
	//var row4 = [] float64 { 0.0, 0.0, 0.0, 0.0, 0.5, 0.5}
	//var row5 = [] float64 { 0.0, 0.0, 0.0, 0.5, 0.0, 0.5}
	//var row6 = [] float64 { 0.0, 0.0, 0.0, 1.0, 0.0, 0.0}

	tm := make ( [][] float64, rows)
	for i := range tm {
		tm[i] = make ( [] float64, cols)
	}


	g.Matrix = mm
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
			fmt.Printf( "  m[%d][%d] = %.2f ", i,j,tm[i][j])
		}
		fmt.Println("")
	}

}
