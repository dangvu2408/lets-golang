package main

import (
	"fmt"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func add(a, b string) string {
	a = reverse(a)
	b = reverse(b)

	if len(a) < len(b) {
		a, b = b, a
	}

	result := ""
	carry := 0

	for i := 0; i < len(a); i++ {
		digitA := int(a[i] - '0')
		digitB := 0
		if i < len(b) {
			digitB = int(b[i] - '0')
		}

		sum := digitA + digitB + carry
		carry = sum / 10
		result += string('0' + sum%10)
	}

	if carry > 0 {
		result += "1"
	}

	return reverse(result)
}

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var a, b string
		fmt.Scan(&a, &b)
		fmt.Println(add(a, b))
	}
}
