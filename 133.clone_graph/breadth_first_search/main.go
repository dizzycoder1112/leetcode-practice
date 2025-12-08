package main

import (
	"fmt"

	"clone_graph"
)

func main() {
	testCases := testcase.GetTestCases()
	for i, adjList := range testCases {
		graph := testcase.BuildGraph(adjList)
		cloned := cloneGraph(graph)
		result := testcase.GraphToAdjList(cloned)
		fmt.Printf("Test Case %d: %v\n", i+1, result)
	}
}

func cloneGraph(node *testcase.Node) *testcase.Node {
	if node == nil {
		return nil
	}

	// visited: original node -> cloned node
	visited := make(map[*testcase.Node]*testcase.Node)

	// Clone the starting node first
	visited[node] = &testcase.Node{Val: node.Val}

	// BFS queue
	queue := []*testcase.Node{node}

	for len(queue) > 0 {
		// Pop from queue
		curr := queue[0]
		queue = queue[1:]

		// Traverse all neighbors
		for _, neighbor := range curr.Neighbors {
			// If neighbor hasn't been cloned yet
			if _, ok := visited[neighbor]; !ok {
				// Clone the neighbor
				visited[neighbor] = &testcase.Node{Val: neighbor.Val}
				// Add to queue for processing
				queue = append(queue, neighbor)
			}
			// Connect cloned nodes
			visited[curr].Neighbors = append(visited[curr].Neighbors, visited[neighbor])
		}
	}

	return visited[node]
}
