package _2_深度优先搜索

import "fmt"

var result8 []int
var count = 8

func init() {
	result8 = make([]int, count)
}


func Cal8Queens(row int) {
	if row >= count {
		printQueens(result8)
		return
	}
	for column := 0; column < count; column ++ {
		if isOk(row, column) {
			result8[row] = column
			Cal8Queens(row + 1)
		}
	}
}
func isOk(row, column int) bool {//判断row行column列放置是否合适
	leftUp, rightUp := column + 1, column + 1
	for i := row-1; i >= 0;i -- {
		if result8[i] == column {
			return false
		}
		if leftUp >= 0 {
			if result8[i] == leftUp {
				return false
			}
		}
		if rightUp < 8 {
			if result8[i] == rightUp {
				return false
			}
		}
		leftUp--
		rightUp++
	}
	return true
}
//打印8*8矩阵
func printQueens(result []int) {
	for row := 0; row < count; row ++ {
		for column := 0; column < count; column ++ {
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}