package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const N = 5

var columns = "BINGO"

func randomInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func columnIndex(c byte) int {
	return strings.IndexByte(columns, c)
}

func printBoard(board [N][N]int, selected [N][N]bool) {
	fmt.Println("\n    B   I   N   G   O")
	for i := 0; i < N; i++ {
		fmt.Printf(" %d ", i+1)
		for j := 0; j < N; j++ {
			if i == 2 && j == 2 {
				fmt.Printf(" *  ")
			} else if selected[i][j] {
				fmt.Printf("[%2d]", board[i][j])
			} else {
				fmt.Printf(" %2d ", board[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var board [N][N]int
	var selected [N][N]bool
	var lineSums [12]int
	var lineCompleted [12]bool
	linesCleared := 0

	for col := 0; col < N; col++ {
		used := make(map[int]bool)
		for row := 0; row < N; row++ {
			if row == 2 && col == 2 {
				board[row][col] = 0
				continue
			}
			var num int
			for {
				num = randomInRange(col*15+1, col*15+15)
				if !used[num] {
					used[num] = true
					break
				}
			}
			board[row][col] = num
		}
	}
	selected[2][2] = true

	for i := 0; i < 5; i++ {
		rowSum := 0
		colSum := 0
		for j := 0; j < 5; j++ {
			rowSum += board[i][j]
			colSum += board[j][i]
		}
		lineSums[i] = rowSum
		lineSums[i+5] = colSum
	}
	for i := 0; i < 5; i++ {
		lineSums[10] += board[i][i]
		lineSums[11] += board[i][4-i]
	}
	minSum := lineSums[0]
	for _, val := range lineSums {
		if val < minSum {
			minSum = val
		}
	}

	for linesCleared < 5 {
		printBoard(board, selected)

		var inputCol string
		var inputRow int
		fmt.Print("\nNhập ô (VD: B 2), hoặc 'exit' để thoát: ")
		fmt.Scan(&inputCol)

		if strings.ToLower(inputCol) == "exit" {
			break
		}
		fmt.Scan(&inputRow)

		col := columnIndex(inputCol[0])
		row := inputRow - 1

		if col < 0 || col >= N || row < 0 || row >= N {
			fmt.Println("Nhập sai. Vui lòng thử lại.")
			continue
		}

		if selected[row][col] {
			fmt.Println("Ô này đã được chọn.")
			continue
		}

		selected[row][col] = true

		for i := 0; i < 5; i++ {
			if lineCompleted[i] {
				continue
			}
			full := true
			for j := 0; j < 5; j++ {
				if !selected[i][j] {
					full = false
				}
			}
			if full {
				lineCompleted[i] = true
				linesCleared++
				if lineSums[i] == minSum {
					fmt.Printf("Bạn đã hoàn thành hàng %d với tổng nhỏ nhất!\n", i+1)
				} else {
					fmt.Printf("Bạn đã hoàn thành hàng %d, nhưng không phải tổng nhỏ nhất.\n", i+1)
				}
			}
		}

		for j := 0; j < 5; j++ {
			if lineCompleted[5+j] {
				continue
			}
			full := true
			for i := 0; i < 5; i++ {
				if !selected[i][j] {
					full = false
				}
			}
			if full {
				lineCompleted[5+j] = true
				linesCleared++
				if lineSums[5+j] == minSum {
					fmt.Printf("Bạn đã hoàn thành cột %c với tổng nhỏ nhất!\n", columns[j])
				} else {
					fmt.Printf("Bạn đã hoàn thành cột %c, nhưng không phải tổng nhỏ nhất.\n", columns[j])
				}
			}
		}

		if !lineCompleted[10] {
			full := true
			for i := 0; i < 5; i++ {
				if !selected[i][i] {
					full = false
				}
			}
			if full {
				lineCompleted[10] = true
				linesCleared++
				if lineSums[10] == minSum {
					fmt.Println("Bạn đã hoàn thành đường chéo chính với tổng nhỏ nhất!")
				} else {
					fmt.Println("Bạn đã hoàn thành đường chéo chính, nhưng không phải tổng nhỏ nhất.")
				}
			}
		}

		if !lineCompleted[11] {
			full := true
			for i := 0; i < 5; i++ {
				if !selected[i][4-i] {
					full = false
				}
			}
			if full {
				lineCompleted[11] = true
				linesCleared++
				if lineSums[11] == minSum {
					fmt.Println("Bạn đã hoàn thành đường chéo phụ với tổng nhỏ nhất!")
				} else {
					fmt.Println("Bạn đã hoàn thành đường chéo phụ, nhưng không phải tổng nhỏ nhất.")
				}
			}
		}
	}

	if linesCleared >= 5 {
		fmt.Println("\nBạn đã hoàn thành 5 dòng! Chúc mừng phá đảo trò chơi!")
	} else {
		fmt.Println("\nBạn đã thoát trò chơi.")
	}
}
