package main

import "fmt"

func (g Graph) performTopologicalSort(s int, visited []bool, stack *Stack) {
	//set the current vertice to visited
	visited[s] = true
	//loop through its adjacencent vertices
	for _, adj := range g.Adj[s] {
		//if the vertice has not been visited then
		if !visited[adj] {
			//recursively perform DFS on it
			g.performTopologicalSort(adj, visited, stack)
		}
	}
	//push the number to the stack
	stack.push(s)
}

func (g Graph) topologicalSort() {
	stack := Stack{}
	//create a list of visited vertices
	visited := make([]bool, g.V)
	//Loop through the vertices
	for i := 0; i < g.V; i++ {
		//if the vertice is not visited then
		if !visited[i] {
			//topologically sort it
			g.performTopologicalSort(i, visited, &stack)
		}
	}

	//loop through the stack to print the sort
	for !stack.empty() {
		fmt.Print(stack.pop(), " ")
	}
}
