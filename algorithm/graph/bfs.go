package main

import (
	"fmt"
)

type Graph map[int][]int

func bfs(graph Graph, start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
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

	fmt.Print("BFS từ đỉnh 0: ")
	bfs(graph, 0)
}
