package main

import (
	"fmt"
	"math"
)

const (
	numPlayers = 5
)

func main() {
	points := make([]int, numPlayers)
	for i := range points {
		points[i] = 10
	}

	eliminated := 0

	for eliminated < numPlayers-1 {
		fmt.Println("\n--- Vòng chơi mới ---")
		choices := make([]int, numPlayers)
		activePlayers := 0

		for i := 0; i < numPlayers; i++ {
			if points[i] > 0 {
				fmt.Printf("Người chơi %d (điểm %d), nhập số (0-100): ", i+1, points[i])
				fmt.Scan(&choices[i])
				activePlayers++
			} else {
				choices[i] = -1
			}
		}

		zeroIdx, hundredIdx := -1, -1
		for i, val := range choices {
			if val == 0 {
				zeroIdx = i
			} else if val == 100 {
				hundredIdx = i
			}
		}
		if eliminated >= 3 && zeroIdx != -1 && hundredIdx != -1 && points[hundredIdx] > 0 {
			fmt.Printf("Người chơi %d chọn 100 thắng theo luật 3!\n", hundredIdx+1)
			for i := 0; i < numPlayers; i++ {
				if i != hundredIdx && points[i] > 0 {
					points[i]--
					if points[i] == 0 {
						eliminated++
					}
				}
			}
			continue
		}

		sum := 0.0
		count := 0
		for _, val := range choices {
			if val != -1 {
				sum += float64(val)
				count++
			}
		}
		avg := math.Round(sum / float64(count) * 0.8)
		fmt.Printf("Mục tiêu: %.0f\n", avg)

		bestIdx := -1
		bestDiff := 1e9
		for i, val := range choices {
			if val != -1 {
				diff := math.Abs(float64(val) - avg)
				if diff < bestDiff {
					bestDiff = diff
					bestIdx = i
				}
			}
		}

		if eliminated >= 1 {
			countMap := make(map[int][]int)
			for i, val := range choices {
				if val != -1 {
					countMap[val] = append(countMap[val], i)
				}
			}
			for _, idxs := range countMap {
				if len(idxs) >= 2 {
					for _, i := range idxs {
						points[i]--
						fmt.Printf("Người chơi %d bị trừ điểm do trùng số (luật 1)\n", i+1)
						if points[i] == 0 {
							eliminated++
						}
					}
				}
			}
		}

		if eliminated >= 2 {
			for i, val := range choices {
				if val != -1 && float64(val) == avg {
					fmt.Printf("Người chơi %d đoán đúng mục tiêu (luật 2)!\n", i+1)
					for j := 0; j < numPlayers; j++ {
						if j != i && points[j] > 0 {
							points[j] -= 2
							if points[j] < 0 {
								points[j] = 0
							}
							if points[j] == 0 {
								eliminated++
							}
						}
					}
					goto endRound
				}
			}
		}

		for i := 0; i < numPlayers; i++ {
			if i != bestIdx && points[i] > 0 {
				points[i]--
				if points[i] == 0 {
					eliminated++
				}
			}
		}
		fmt.Printf("Người chơi %d thắng vòng này!\n", bestIdx+1)

	endRound:
		fmt.Println("Điểm hiện tại:")
		for i, p := range points {
			fmt.Printf("Người chơi %d: %d điểm\n", i+1, p)
		}
	}

	fmt.Println("\n--- Trò chơi kết thúc ---")
	for i, p := range points {
		if p > 0 {
			fmt.Printf("🏆 Người chơi %d chiến thắng với %d điểm!\n", i+1, p)
		}
	}
}
