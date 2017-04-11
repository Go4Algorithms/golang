package graph

//returns the shortest distance between each pair of nodes
func FloydWarshall(g *Graph)[][]int {
	if g.Weighted {
		return FloydWarshallM(g.Matrix)
	}
	return nil
}

func FloydWarshallM(graph [][] int)[][]int {
	var nrVertices int
	//check if it's a for this purpose valid matrix
	if len(graph[0]) == len(graph) && len(graph)>0 {
		nrVertices = len(graph)
	}else {
		return nil
	}
	distance := make([][]int, nrVertices)
	for i:=0; i<nrVertices;i++ {
		distance[i] = make([]int, nrVertices)
	}
	for i:=0; i<nrVertices;i++ {
		for j:=0; j<nrVertices;j++ {
			// the weighted graph matrix sets non existing edges to 0
			if i!=j && graph[i][j]==0 {
				//don't use max value because when you add to it it will equal math.min
				distance[i][j] = 1000000
			}else{
				distance[i][j] = graph[i][j]
			}
		}
	}
	
	for k:=0; k<nrVertices;k++ {
		for i:=0; i<nrVertices;i++ {
			for j:=0; j<nrVertices;j++ {
				if distance[i][k] + distance[k][j] < distance[i][j]{
                    			distance[i][j] = distance[i][k] + distance[k][j]
				}
			}
		}
	}
	return distance
}
