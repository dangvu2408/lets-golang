package main

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append phần còn lại
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	arr := []int{5, 7, 2, 9, 1, 3, 0, 4, 8, 6}
	arr = mergeSort(arr)
	fmt.Println(arr)
}
