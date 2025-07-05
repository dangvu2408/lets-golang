package main

import "fmt"

func bubble(a []int) {
	n := len(a)
	i := 0
	sorted := false
	for !sorted && i < n-1 {
		sorted = true
		for j := n - 2; j >= i; j-- {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				sorted = false
			}
		}
		i++
	}
}

func main() {
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	bubble(arr)
	fmt.Println(arr)
}
