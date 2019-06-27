package main

import "fmt"

func main() {
	backPack(10, []int{1, 2, 5, 6})
}

func backPack(total int, packs []int) {
	len := len(packs)
	//len2 := total
	var sum [10][50]int
	for i := 0; i < len; i++ {
		sum[i][0] = 0
	}
	for j := 0; j < total+1; j++ {
		if j < packs[0] {
			sum[0][j] = packs[0]
		}
	}
	for i := 1; i < len; i++ {
		for j := 1; j < total+1; j++ {
			if j >= packs[i] {
				sum[i][j] = max(sum[i-1][j], sum[i-1][j-packs[i]]+packs[i])
				sum[i][j] = max(sum[i-1][j], sum[i-1][j-packs[i]]+packs[i])
			} else {
				sum[i][j] = sum[i-1][j]
			}
		}
	}
	fmt.Println(sum)
	best := sum[len-1][total]
	fmt.Println(best)
}
func max(a, b int) int {
	if a > b {
		return b
	}
	return a
}
