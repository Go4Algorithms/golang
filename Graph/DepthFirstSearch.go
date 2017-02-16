package graph

import(
    "p_SimpleStack"
    "math"
    )
    
//function that returns the depth of each node
func DepthFirstSearch(graph [][]int, source int)([]int){
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
    stack = p_SimpleStack.New()
    return dfs(graph,source,distances,visited)
}

func dfs(graph[][]int, node int, distances[]int, visited[]bool, stack SimpleStack)([]int) {
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
    if stack.isEmpty(){
        return distances
    }
    if !foundNew {
        i = stack.Pop()
    }
    return dfs(graph,i,distances,visited,stack)
}