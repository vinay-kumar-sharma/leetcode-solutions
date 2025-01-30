// 2493. Divide Nodes Into the Maximum Number of Groups
package main

import (
	"fmt"
)

func magnificentSets(n int, edges [][]int) int {
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]int, n)
	components := make([][]int, 0)

	// Find connected components
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			queue := []int{i}
			component := []int{}
			visited[i] = 1
			isBipartite := true

			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]
				component = append(component, node)

				for _, neighbor := range graph[node] {
					if visited[neighbor] == 0 {
						visited[neighbor] = -visited[node]
						queue = append(queue, neighbor)
					} else if visited[neighbor] == visited[node] {
						isBipartite = false
					}
				}
			}

			if !isBipartite {
				return -1
			}
			components = append(components, component)
		}
	}

	// Compute max depth using BFS
	maxGroups := 0
	for _, component := range components {
		maxDepth := 0
		for _, node := range component {
			maxDepth = max(maxDepth, bfs(graph, node, n))
		}
		maxGroups += maxDepth
	}

	return maxGroups
}

func bfs(graph [][]int, start, n int) int {
	visited := make([]int, n)
	queue := []int{start}
	visited[start] = 1
	maxDepth := 1

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				if visited[neighbor] == 0 {
					visited[neighbor] = visited[node] + 1
					maxDepth = max(maxDepth, visited[neighbor])
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return maxDepth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test the function
func main() {
	n := 6
	edges := [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}
	fmt.Println(magnificentSets(n, edges)) // Expected Output: 4
}