package main

import "fmt"

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
