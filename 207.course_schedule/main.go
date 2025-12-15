package main

import "fmt"

func main() {
	// Test case 1: Basic case - can finish
	// 0 → 1 (take 0 first, then 1)
	fmt.Println("Test 1:", canFinish(2, [][]int{{1, 0}})) // Expected: true

	// Test case 2: Cycle - cannot finish
	// 0 → 1, 1 → 0 (deadlock!)
	fmt.Println("Test 2:", canFinish(2, [][]int{{1, 0}, {0, 1}})) // Expected: false

	// Test case 3: No prerequisites
	fmt.Println("Test 3:", canFinish(3, [][]int{})) // Expected: true

	// Test case 4: Multiple dependencies
	// 0 → 1, 0 → 2, 1 → 2
	fmt.Println("Test 4:", canFinish(3, [][]int{{1, 0}, {2, 0}, {2, 1}})) // Expected: true

	// Test case 5: Longer cycle
	// 0 → 1 → 2 → 0
	fmt.Println("Test 5:", canFinish(3, [][]int{{1, 0}, {2, 1}, {0, 2}})) // Expected: false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. Build graph and in-degree array
	graph := make([][]int, numCourses)  // adjacency list
	inDegree := make([]int, numCourses) // count of prerequisites

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		graph[prereq] = append(graph[prereq], course) // prereq → course
		inDegree[course]++
	}

	// 2. Queue starts with courses that have NO prerequisites (in-degree = 0)
	queue := []int{}
	for i := range numCourses {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. Process: each time we "take" a course, reduce in-degree of dependent courses
	completed := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		completed++

		for _, next := range graph[current] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 4. If we completed all courses, no cycle!
	return completed == numCourses
}
