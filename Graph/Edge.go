package graph

type Edge struct {
	Name string
	ID int
	Directed bool
	From *Node
	To *Node
	Weight int
}

//edge with weight
func newWeightedEdge(from *Node, to *Node, directed bool, argName string, weight int)*Edge {
    var r = new (Edge)
	r.Name = argName
	r.Directed = directed
	r.From = from
	r.To = to
	r.Weight = weight
	return r
}

// newEdge ( from, to, argName)
func newEdge ( from *Node, to *Node, directed bool, argName string) *Edge {
	var r = new (Edge)
	r.Name = argName
	r.From = from
	r.To = to
	return r
}

