package main

import "fmt"

func insertion(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		tmp := a[i]
		j := i - 1
		for j >= 0 && tmp < a[j] {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = tmp
	}
}

func main() {
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	insertion(arr)
	fmt.Println(arr)
}
