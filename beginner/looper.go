package main

import "fmt"

func main() {
	for i, j := 0, 10; i < 10 && j > 0; i, j = i+1, j-1 {
		fmt.Printf("%d - %d\n", i, j)
	}

}
