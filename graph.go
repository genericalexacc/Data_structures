package main

//Graph structure
type Graph struct {
	V   int
	Adj [][]int
}

//MakeGraph graph constructor
func MakeGraph(V int) Graph {
	return Graph{
		//number of vertices in the graph
		V,
		//adjacency list is a 2d array of ints
		make([][]int, V),
	}
}

//Add an edge to the graph
func (g Graph) addEdge(v, w int) {
	g.Adj[v] = append(g.Adj[v], w)
}
