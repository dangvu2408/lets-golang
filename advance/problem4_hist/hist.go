package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func largest(hist []int) int {
	stack := []int{}
	maxArea := 0
	hist = append(hist, 0)

	for i := 0; i < len(hist); i++ {
		for len(stack) > 0 && hist[i] < hist[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := hist[top]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "0" {
			break
		}

		parts := strings.Fields(line)
		n, _ := strconv.Atoi(parts[0])
		if n == 0 {
			continue
		}

		hist := make([]int, n)
		for i := 0; i < n; i++ {
			hist[i], _ = strconv.Atoi(parts[i+1])
		}

		fmt.Println(largest(hist))
	}
}
