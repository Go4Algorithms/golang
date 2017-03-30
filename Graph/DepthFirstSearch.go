package graph

import(
    "github.com/Go4Algorithms/golang/DataStructures/p_SimpleStack"
    "math"
    )
    
func DepthFirstSearch(g *Graph, source int)([]int){
	if g.Weighted {
		return DepthFirstSearchM(g.Matrix, source)
	}
	return nil
}
//function that returns the depth of each node
func DepthFirstSearchM(graph [][]int, source int)([]int){
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
    stack := p_SimpleStack.New()
    return dfs(graph,source,distances,visited,stack)
}

func dfs(graph[][]int, node int, distances[]int, visited[]bool, stack p_SimpleStack.SimpleStack)([]int) {
    foundNew := false
    var i int
    for i=0; i<len(graph); i++ {
        if graph[node][i]!=0&&!visited[i] {
            foundNew = true
            visited[i] = true
            distances[i] = distances[node]+1
            stack.Push(i)
            break
        }
    }
    if stack.IsEmpty(){
        return distances
    }
    if !foundNew {
        i = stack.Pop()
    }
    return dfs(graph,i,distances,visited,stack)
}
