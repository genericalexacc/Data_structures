package main

import "fmt"

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

//BFS performs a breadth first search over the graph
func (g Graph) BFS(s int) {
	//we create an array of bools with indexes that we have visited
	visited := make([]bool, g.V)
	//we create a queue
	queue := Queue{[]int{}, 0}
	//we set the current vertice as visited
	visited[s] = true
	//we append the current vertice to the queue
	queue.enqueue(s)
	//we loop until the queue is empty
	for !queue.empty() {
		//set the current vertice to the first item in the queue
		s = queue.dequeue()
		//print it out
		fmt.Print(s, " ")
		//loop through the adjacent vertices
		for _, adj := range g.Adj[s] {
			//if the vertice has not been visited yet then
			if !visited[adj] {
				//set the vertice to visited
				visited[adj] = true
				//append it to the queue
				queue.enqueue(adj)
			}
		}
	}
}
