package main

import (
	"fmt"
	"math"
)

func (wg *WeightedGraph) minDistance(dist []int, sptSet []bool) int {
	//initialize min int to the maximum possible value
	min := math.MaxInt64
	//min index is 0
	var minIndex int

	//We go through the vertices we have
	for v := 0; v < wg.V; v++ {
		//if the shortest path tree hasn't been done yet and the minimum distance is smaller then
		if sptSet[v] == false && dist[v] <= min {
			//we set the minimum to the shortest path
			min = dist[v]
			//and set the minimum index to the current vertice
			minIndex = v
		}
	}

	//return the minimum index
	return minIndex
}

func (wg *WeightedGraph) dijkstra(src int) {
	//Create a distance array of length (# vertices)
	dist := make([]int, wg.V)
	//Create an array that keeps track of whether we did its shortest path tree of length (# vertices)
	sptSet := make([]bool, wg.V)
	//set all distances to max value
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt64
	}
	//set the distance to self to 0
	dist[src] = 0
	//Go through all vertices
	for count := 0; count < wg.V-1; count++ {
		//get the verticeIndex with minimum distance
		u := wg.minDistance(dist, sptSet)
		//set its index as done in the shortest path tree array
		sptSet[u] = true
		//go through all the vertices we have
		for v := 0; v < wg.V; v++ {
			//if its not done yet, its distance is not 0, its distance is not maximum and
			//distance to u + distance from u to v is smaller than distance to v then
			if !sptSet[v] && wg.Adj[u][v] != 0 && dist[u] != math.MaxInt64 && dist[u]+wg.Adj[u][v] < dist[v] {
				//set the distance to sum of distance to u + distance from u to v
				dist[v] = dist[u] + wg.Adj[u][v]
			}
		}
	}

	//Itterate through distances
	for i := 0; i < len(dist); i++ {
		// Print out the results
		fmt.Println("Distance from", src, "->", i, "is", dist[i])
	}
}
