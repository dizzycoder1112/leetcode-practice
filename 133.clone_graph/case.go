package testcase

type Node struct {
	Val       int
	Neighbors []*Node
}

// GetTestCases returns test cases as adjacency lists
// Each test case: adjList[i] contains neighbors of node i+1
func GetTestCases() [][][]int {
	return [][][]int{
		{{2, 4}, {1, 3}, {2, 4}, {1, 3}}, // 4 nodes graph, Expected: same structure
		{{}},                              // Single node, no neighbors
		{},                                // Empty graph
	}
}

// BuildGraph builds a graph from adjacency list
func BuildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	// Create all nodes first
	nodes := make([]*Node, len(adjList))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}

	// Connect neighbors
	for i, neighbors := range adjList {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}

	return nodes[0]
}

// GraphToAdjList converts graph back to adjacency list for verification
func GraphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	visited := make(map[*Node]bool)
	nodes := []*Node{}

	// BFS to collect all nodes
	queue := []*Node{node}
	visited[node] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		nodes = append(nodes, curr)

		for _, neighbor := range curr.Neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	// Sort by Val and build adjacency list
	adjList := make([][]int, len(nodes))
	for _, n := range nodes {
		for _, neighbor := range n.Neighbors {
			adjList[n.Val-1] = append(adjList[n.Val-1], neighbor.Val)
		}
	}

	return adjList
}
