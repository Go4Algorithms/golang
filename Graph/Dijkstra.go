package graph
//dijkstra's single shortest path algorithm
import(
	"math"
)

//function that returns the edges used for the shortest path
func DijkstraEdges(g *Graph, source *Node)([]Edge){
	if !g.Weighted {
		return nil
	}
    var nrVertices,src int
    nrVertices = len(g.Nodes)
    for i:=0; i<nrVertices; i++ {
        if *source == g.Nodes[i] {
            src = i
        }
    }
    fromNode := make([]int, len(g.Edges))
    pickedUpEdges := make([]Edge, nrVertices-1)
    distance := make([]int,nrVertices)
	includedSet := make([]bool,nrVertices)
	graph := g.Matrix
    for i:=0; i<nrVertices;i++ {
		distance[i] = 1000000
		includedSet[i] = false
	}
	//distance source to itself is 0, so it gets picked up first
	distance[src] = 0
	//find shortest path for all vertices
	for count:=-1; count<nrVertices-1;count++{
		u := minDistance(distance,includedSet,nrVertices)
		includedSet[u] = true
		for i:= 0; i<len(g.Edges);i++ {
		    if *g.Edges[i].From == g.Nodes[fromNode[u]] && *g.Edges[i].To == g.Nodes[u] {
		        pickedUpEdges[count] = g.Edges[i]
		    }
		}
		for v:=0;v<nrVertices;v++ {
			if !includedSet[v]&&graph[u][v]!=0&&distance[u]+graph[u][v]<distance[v] {
				distance[v] = distance[u]+graph[u][v]
				fromNode[v] = u
			}
		}
	}
	return pickedUpEdges
}

//A function to find the not included vertex with the minimum distance 
func minDistance( distance[]int,includedSet[]bool, nrVertices int)(int) {
	min := math.MaxInt32
	minIndex := -1
	for v:=0; v<nrVertices; v++{
		if !includedSet[v] && distance[v] <= min {
			min = distance[v]
			minIndex = v
		}
	}
	return minIndex
}

func DijkstraDistance(g *Graph, src int)([]int){
	if g.Weighted {
		return DijkstraDistanceM(g.Matrix, src)
	}
	return nil
}
//this function returns the distances between each vertex and the source
func DijkstraDistanceM(graph[][]int,src int)([]int) {
	var nrVertices int
	//check if it's a for this purpose valid matrix
	if len(graph[0]) == len(graph) && len(graph)>0 {
		nrVertices = len(graph)
	} else {
		return nil
	}
	distance := make([]int,nrVertices)
	includedSet := make([]bool,nrVertices)
	for i:=0; i<nrVertices;i++ {
		distance[i] = math.MaxInt32
		includedSet[i] = false
	}
	//distance source to itself is 0, so it gets picked up first
	distance[src] = 0
	//find shortest path for all vertices
	for count:=0; count<nrVertices-1;count++{
		u := minDistance(distance,includedSet,nrVertices)
		includedSet[u] = true
		for v:=0;v<nrVertices;v++ {
			if !includedSet[v]&&graph[u][v]!=0&&distance[u]!= math.MaxInt32&&distance[u]+graph[u][v]<distance[v] {
				distance[v] = distance[u]+graph[u][v]
			}
		}
	}
	return distance	
}

