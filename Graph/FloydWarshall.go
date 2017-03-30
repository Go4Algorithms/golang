package graph
import(
	"math"
)
//returns the shortest distance between each pair of nodes
func FloydWarshall(g *Graph)[][]int {
	if g.Weighted {
		return FloydWarshallM(g.Matrix)
	}
	return nil
}

func FloydWarshallM(graph [][]Graph)[][]int {
	nrVertices := len(graph)
	distance := make([][]int, nrVertices)
	for i := 0; i < nrVertices; i++ {
		for j := 0; j < nrVertices; j++ {
			// the weighted graph matrix sets non existing edges to 0
			if i != j && graph[i][j] == 0 {
				distance[i][j] = math.MaxInt32
			}else{
				distance[i][j] = graph[i][j]
			}
		}
	}

	for k := 0; k<nrVertices; k++ {
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
