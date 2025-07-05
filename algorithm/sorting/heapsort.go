package main

import "fmt"

func makeSubHeap(a []int, r int, n int) {
	if 2*r+1 >= n {
		return
	}
	c := 2*r + 1
	if c+1 < n && a[c+1] > a[c] {
		c = c + 1
	}

	if a[r] < a[c] {
		a[r], a[c] = a[c], a[r]
		makeSubHeap(a, c, n)
	}
}

func makeHeap(a []int, n int) {
	if n < 2 {
		return
	}
	var m int
	m = n/2 - 1
	for i := m; i >= 0; i-- {
		makeSubHeap(a, i, n)
	}
}

func heap(a []int, n int) {
	if n < 2 {
		return
	}
	makeHeap(a, n)
	for i := 1; i < n; i++ {
		a[0], a[n-i] = a[n-i], a[0]
		makeSubHeap(a, 0, n-i)
	}
}

func main() {
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	heap(arr, len(arr))
	fmt.Println(arr)
}
