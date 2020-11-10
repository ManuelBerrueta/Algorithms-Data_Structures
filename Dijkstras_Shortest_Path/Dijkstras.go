package main

import (
	"fmt"
)

type Edge struct {
	neighbor string
	weight   float64
}

// Vertex represents a node or vertex on a graph
type Vertex struct {
	ID               string
	visited          bool
	shortestDistance float64
	previousVertex   string   // Set
	neighbors        []string // Store the connections
	edges            []Edge
}

// Graph represents a group of vetices or nodes
type Graph struct {
	vertices []Vertex
}

// createUnvisitedQueue returns a list ([]string) of unvisited nodes.
func createUnvisitedQueue() {

}

func initializeGraph() *Graph {
}

func main() {
	fmt.Println("-=0=☼=0=☼=☼=0=☼=☼=☼=[ Dijkstra’s Shortest Path Algorithm ]=☼=☼=☼=0=☼=☼=0=☼=0=-")
	//! Destructions:
	//! Dijkstra's algorithm is used to find the shortest path from a starting node to every other node in the graph from that starting node
	//? We need 2 lists, one to keep track of visited nodes and one of not-visited nodes
	//? Initialize the shortest distance from start vertex to any vertex to infinity, including itself
	//? Starting with starting node A, we will find that the distance from A to itself is 0, and we update that distance and previous vertex
	//? We add the current vertex, to the list of visited vertices.
	//? Next we visit the next vertex with the smallest known distance from the starting node A
	//? From this next vertex B we also look at it's unvisited neighbors
	//? From this vertex we calculate the distance from the starting vertex A to this vertex's (B) neighbors, this distance is known from
	//?  our input updated in the table
	//? If the distance is shorter than the current distance, we update the distance and previous vertex

	visited := []string{}
	notVisited := []string{}
	visited = append(visited, "you", "got")
	fmt.Println(visited)
	fmt.Println(notVisited)

	//* Reference: https://youtu.be/pVfj6mxhdMw, https://youtu.be/_lHSawdgXpI
}
