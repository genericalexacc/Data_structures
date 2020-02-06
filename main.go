package main

func main() {
	wg := MakeWeightedGraph(9)

	wg.addEdge(0, []int{0, 4, 0, 0, 0, 0, 0, 8, 0})
	wg.addEdge(1, []int{4, 0, 8, 0, 0, 0, 0, 11, 0})
	wg.addEdge(2, []int{0, 8, 0, 7, 0, 4, 0, 0, 2})
	wg.addEdge(3, []int{0, 0, 7, 0, 9, 14, 0, 0, 0})
	wg.addEdge(4, []int{0, 0, 0, 9, 0, 10, 0, 0, 0})
	wg.addEdge(5, []int{0, 0, 4, 14, 10, 0, 2, 0, 0})
	wg.addEdge(6, []int{0, 0, 0, 0, 0, 2, 0, 1, 6})
	wg.addEdge(7, []int{8, 11, 0, 0, 0, 0, 1, 0, 7})
	wg.addEdge(8, []int{0, 0, 2, 0, 0, 0, 6, 7, 0})

	wg.dijkstra(0)
}
