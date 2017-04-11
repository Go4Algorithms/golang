package graph

import (
	"github.com/Go4Algorithms/golang/DataStructures/p_SimpleQueue"
	"math"
)
    
func BreadthFirstSearch(g *Graph, source int)[]int{
	if g.Weighted {
		return BreadthFirstSearchM(g.Matrix,source)
	}
	return nil
}
//Breadth First Search Algorithms
//function that returns the distance to each Vertex
func BreadthFirstSearchM(graph[][]int,source int)[]int{
    var nrVertices int
	//check if it's a for this purpose valid matrix
	if len(graph[0]) == len(graph) && len(graph)>0 {
		nrVertices = len(graph)
	} else {
		return nil
	}
    distances := make([]int,nrVertices)
    visited := make([]bool,nrVertices)
    for i:=0; i<nrVertices; i++ {
        distances[i] = math.MaxInt32
        visited[i] = false
    }
	distances[source]=0
	visited[source]=true
    queue := p_SimpleQueue.New()
    return bfs(graph,source,distances,visited,queue)
}

func bfs(graph[][]int,node int,distances[]int,visited[]bool,queue p_SimpleQueue.SimpleQueue)[]int{
    for i:=0; i<len(graph); i++ {
        if graph[node][i]!=0 && !visited[i] {
            visited[i] = true
            distances[i] = distances[node]+1
            queue.Push(i)
        }
    }
    if queue.IsEmpty() {
        return distances
    }
    return bfs(graph,queue.Pop(),distances,visited,queue)
}
