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
	return dfs(node, visited)
}

func dfs(node *testcase.Node, visited map[*testcase.Node]*testcase.Node) *testcase.Node {
	// If already cloned, return the clone
	if clone, ok := visited[node]; ok {
		return clone
	}

	// Clone current node
	clone := &testcase.Node{Val: node.Val}
	visited[node] = clone

	// Recursively clone all neighbors
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(neighbor, visited))
	}

	return clone
}
