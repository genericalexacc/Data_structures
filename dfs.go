package main

import "fmt"

func (g Graph) performDFS(s int, visited []bool) {
	//set the current vertice to visited
	visited[s] = true
	//print it out
	fmt.Print(s, " ")
	//loop through its adjacencent vertices
	for _, adj := range g.Adj[s] {
		//if the vertice has not been visited then
		if !visited[adj] {
			//recursively perform DFS on it
			g.performDFS(adj, visited)
		}
	}
}

//DFS performs a depth first search over the graph
func (g Graph) DFS(s int) {
	//create a list of visited vertices
	visited := make([]bool, g.V)
	//perform depth first search recursively
	g.performDFS(s, visited)
}
