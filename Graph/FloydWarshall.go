package graph
import(
	"math"
)
//returns the shortest distance between each pair of nodes
func floydWarshall(graph[][]int)[][]int {
	var nrVertices int
	//check if it's a for this purpose valid matrix
	if len(graph[0]) == len(graph) {
		nrVertices = len(graph)
	} else {
		return nil
	}
	var distance = make([][]int, nrVertices)
	for i:=0; i<nrVertices;i++ {
		for j:=0; j<nrVertices;j++ {
			// the weighted graph matrix sets non existing edges to 0
			if i!=j && graph[i][j]==0 {
				distance[i][j] = math.MaxInt32
			}
			else{
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
