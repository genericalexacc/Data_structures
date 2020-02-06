package main

//WeightedGraph structure
type WeightedGraph struct {
	V   int
	Adj [][]int
}

//MakeWeightedGraph WeightedGraph constructor
func MakeWeightedGraph(V int) WeightedGraph {
	//Create array of length V filled with 0s
	line := make([]int, V)
	//Create array of arrays of length V
	prePop := make([][]int, V)
	//Insert the arrays so that we have an array of arrays filled with 0s
	for i := 0; i < len(prePop); i++ {
		prePop[i] = line
	}
	//return the created weighted graph
	return WeightedGraph{
		V,
		prePop,
	}
}

//Add an edge to the graph
func (wg WeightedGraph) addEdge(v int, w []int) {
	//Add the edge itself
	wg.Adj[v] = w
	//Go through the previous edges and modify their distances to the new vertice
	for index, weight := range w {
		wg.Adj[index][v] = weight
	}
}
