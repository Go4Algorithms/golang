package graph

type Edge struct {
	Name string
	ID int
	From *Node
	To *Node
	Weight int
}

//edge with weight
func NewWeightedEdge(from *Node, to *Node, argName string, weight int, id int)*Edge {
    var r = new (Edge)
	r.Name = argName
	r.From = from
	r.To = to
	r.Weight = weight
	r.ID = id
	return r
}

// newEdge ( from, to, argName)
func NewEdge ( from *Node, to *Node, argName string, id int) *Edge {
	var r = new (Edge)
	r.Name = argName
	r.From = from
	r.To = to
	r.ID = id
	return r
}

