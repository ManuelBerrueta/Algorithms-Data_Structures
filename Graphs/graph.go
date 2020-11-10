package main

// Reference: https://flaviocopes.com/golang-data-structure-graph/

import (
	"fmt"
	"sync"
)

// Node represents a node or vertex in a graph
type Node struct {
	ID string
}

// Graph is the representation of an actual graph including edges
type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// AddNode appends the inNode to the graph
func (graph *Graph) AddNode(inNode *Node) {
	graph.lock.Lock()
	graph.nodes = append(graph.nodes, inNode)
	graph.lock.Unlock()
}

// AddEdge appends an edge (a connection between two nodes) to the graph
func (graph *Graph) AddEdge(node1 *Node, node2 *Node, weight float64) {
	// Locks access to the Graph data structure
	graph.lock.Lock()

	// If the graph is empty
	if graph.edges == nil {
		// Initialize graph
		graph.edges = make(map[Node][]*Node)
	}
	// Else add the incoming edge bidirectionally
	graph.edges[*node1] = append(graph.edges[*node1], node2)
	graph.edges[*node2] = append(graph.edges[*node2], node1)
	graph.lock.Unlock()
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.ID)
}

func main() {
	fmt.Println("Graphs!")
}
