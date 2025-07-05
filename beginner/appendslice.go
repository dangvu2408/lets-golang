package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice co do dai", len(slice), "va suc chua la", cap(slice))

	slice = append(slice, 6)
	fmt.Println("slice sau thay doi co do dai la", len(slice), "va suc chua la", cap(slice))
	fmt.Println(slice)
}
