package main

import (
	"fmt"
)

type Graph map[int][]int

func dfs(graph Graph, start int, visited map[int]bool) {
	if visited[start] {
		return
	}

	visited[start] = true
	fmt.Printf("%d ", start)

	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited)
		}
	}
}

func main() {
	graph := Graph{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0},
		3: {1},
		4: {1, 5},
		5: {4},
	}

	visited := make(map[int]bool)

	fmt.Print("DFS từ đỉnh 0: ")
	dfs(graph, 0, visited)
}
