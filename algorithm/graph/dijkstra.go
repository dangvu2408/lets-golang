package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Cấu trúc cạnh
type Edge struct {
	to     int
	weight int
}

type Graph map[int][]Edge

type Item struct {
	node int
	dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph Graph, start int) map[int]int {
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	visited := make(map[int]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{node: start, dist: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(Item)
		if visited[u.node] {
			continue
		}
		visited[u.node] = true

		for _, edge := range graph[u.node] {
			v := edge.to
			weight := edge.weight
			if dist[u.node]+weight < dist[v] {
				dist[v] = dist[u.node] + weight
				heap.Push(pq, Item{node: v, dist: dist[v]})
			}
		}
	}

	return dist
}

func main() {
	graph := Graph{
		0: {{to: 1, weight: 4}, {to: 2, weight: 1}},
		1: {{to: 3, weight: 1}},
		2: {{to: 1, weight: 2}, {to: 3, weight: 5}},
		3: {},
	}

	start := 0
	distances := Dijkstra(graph, start)

	fmt.Printf("Khoảng cách ngắn nhất từ đỉnh %d:\n", start)
	for node, d := range distances {
		fmt.Printf("Tới %d: %d\n", node, d)
	}
}
