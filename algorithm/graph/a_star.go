package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

type Graph map[int][]Edge

type Node struct {
	id       int
	priority float64
	g        float64
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func heuristic(a, b int) float64 {
	return 0
}

func AStar(graph Graph, start, goal int) ([]int, float64) {
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, Node{id: start, priority: 0, g: 0})

	cameFrom := make(map[int]int)
	gScore := make(map[int]float64)

	for node := range graph {
		gScore[node] = math.Inf(1)
	}
	gScore[start] = 0

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(Node)

		if current.id == goal {
			path := []int{}
			for at := goal; at != start; at = cameFrom[at] {
				path = append([]int{at}, path...)
			}
			path = append([]int{start}, path...)
			return path, gScore[goal]
		}

		for _, edge := range graph[current.id] {
			tentative_g := gScore[current.id] + float64(edge.weight)
			if tentative_g < gScore[edge.to] {
				cameFrom[edge.to] = current.id
				gScore[edge.to] = tentative_g
				f := tentative_g + heuristic(edge.to, goal)
				heap.Push(openSet, Node{id: edge.to, priority: f, g: tentative_g})
			}
		}
	}

	return nil, math.Inf(1)
}

func main() {
	graph := Graph{
		0: {{1, 1}, {2, 4}},
		1: {{2, 2}, {3, 5}},
		2: {{3, 1}},
		3: {},
	}

	start := 0
	goal := 3

	path, cost := AStar(graph, start, goal)
	if path == nil {
		fmt.Println("Không tìm được đường đi.")
	} else {
		fmt.Println("Đường đi ngắn nhất:", path)
		fmt.Println("Chi phí:", cost)
	}
}
