package main

import "fmt"

func main() {
	a := [5]int{8, 4, 3, 7, 5}
	var b []int = a[1:4]
	fmt.Println(b)
	for i, _ := range b {
		b[i]++
	}
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println("Do dai slice", len(b))
	fmt.Println("Do dai toi da slice", cap(b))
	b = a[0:cap(b)]
	fmt.Println("Do dai slice sau sua", len(b))
	fmt.Println("Do dai toi da slice sau sua", cap(b))
}
