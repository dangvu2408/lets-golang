package main

import "fmt"

func selection(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		m := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[m] {
				m = j
			}
		}
		a[m], a[i] = a[i], a[m]
	}
}

func main() {
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	selection(arr)
	fmt.Println(arr)
}
