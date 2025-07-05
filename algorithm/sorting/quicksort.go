package main

import "fmt"

func partition(a []int, first int, last int) {
	if first >= last {
		return
	}
	c := a[first]
	i, j := first+1, last
	for i <= j {
		for i <= j && a[i] <= c {
			i++
		}
		for i <= j && a[j] > c {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[first], a[j] = a[j], a[first]
	partition(a, first, j-1)
	partition(a, j+1, last)
}

func quick(a []int) {
	n := len(a)
	partition(a, 0, n-1)
}

func main() {
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	quick(arr)
	fmt.Println(arr)
}
