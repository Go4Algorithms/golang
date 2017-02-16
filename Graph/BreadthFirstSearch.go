package graph

import(
    "p_SimpleQueue"
    "math"
    )
    
func BreadthFirstSearch(g *Graph, source int)([]int){
	return BreadthFirstSearch(g.Matrix,source)
}
//Breadth First Search Algorithms
//function that returns the distance to each Vertex
func BreadthFirstSearch(graph[][]int,source int)([]int){
    var nrVertices int
	//check if it's a for this purpose valid matrix
	if len(graph[0]) == len(graph) {
		nrVertices = len(graph)
	} else {
		return nil
	}
    var distances = make([]int,nrVertices)
    var visited = make([]bool,nrVertices)
    for i:=0; i<nrVertices; i++ {
        distances[i] = math.MaxInt32
        visited[i] = false
    }
    queue = p_SimpleQueue.New()
    return bfs(graph,source,distances,visited,queue)
}

func bfs(graph[][]int,node int,distances[]int,visited[]bool,queue SimpleQueue)([]int){
    for i:=0; i<len(graph); i++ {
        if graph[node][i]!=0&&!visited[i] {
            visited[i] = true
            distances[i] = distances[node]+1
            queue.Push(i)
        }
    }
    if queue.isEmpty() {
        return distances
    }
    return bfs(graph,queue.Pop(),distances,visited,queue)
}
