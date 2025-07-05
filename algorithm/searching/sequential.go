package main

import (
	"fmt"
	"time"
)

func sequential(a []int, n int, k int) int {
	for i := 0; i < n; i++ {
		if a[i] == k {
			return i
		}
	}
	return -1 // Không tìm thấy
}

func main() {
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	start := time.Now()
	ans := sequential(arr, len(arr), 6)
	elapsed := time.Since(start)
	fmt.Println(ans)
	fmt.Println("Time elapsed:", elapsed.Nanoseconds())
}
