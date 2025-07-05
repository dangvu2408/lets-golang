package main

import (
	"fmt"
	"sort"
	"time"
)

func binary(a []int, b int, e int, k int) int {
	if b > e {
		return -1
	}
	m := (b + e) / 2
	if k == a[m] {
		return m
	} else if k < a[m] {
		return binary(a, b, m-1, k)
	} else {
		return binary(a, m+1, e, k)
	}
}

func binarySearch(k int, a []int, n int) int {
	return binary(a, 0, n-1, k)
}

func main() {
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	sort.Ints(arr)
	start := time.Now()
	ans := binarySearch(6, arr, len(arr))
	elapsed := time.Since(start)
	fmt.Println(ans)
	fmt.Println("Time elapsed:", elapsed.Nanoseconds())
}
