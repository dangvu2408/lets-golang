package main

import "fmt"

func findMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func subArrMax(arr []int, n int) int {
	current := arr[0]
	maximum := arr[0]
	for i := 1; i < n; i++ {
		current = findMax(arr[i], arr[i]+current)
		maximum = findMax(maximum, current)
	}
	return maximum
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Sub Array Max Sum equals:", subArrMax(arr, n))
}
