package main

import "fmt"

func main() {
	a := [5]int{5, 7, 2, 9, 1}
	b := a
	b[0] = 9
	for i, v := range a {
		fmt.Printf("Phan tu thu %d la: %d\n", i, v)
	}

	fmt.Println("Mang b ", b)
}
