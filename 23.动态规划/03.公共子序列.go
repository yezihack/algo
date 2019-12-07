package _3_动态规划

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
)

//最长公共子序列(不是公共子串)
//如fosh,fish 最长公共子串是sh, 公共子序列是fsh
//先抽象成数据模型. 利用矩阵查找关系
//思路: 先目标字符串抽象成字符行, 匹配字符串抽象成字符列. 行与列就成了矩阵
// 若row为行, col为列,比较之间的关系.
// 如果相等则选择左上方单元格的值加1
// 如果不相等则选择左边相邻和上方单元格中较大的那个值.
func SubSequence(a, b string) string {
	rChar := []byte(a)
	cChar := []byte(b)
	//设置哨兵
	rChar = src.NewByte().InsertIndex(rChar, 0, 0)
	cChar = src.NewByte().InsertIndex(cChar, 0, 0)
	iRow, jCol := len(rChar), len(cChar)

	fmt.Println(string(rChar), rChar, string(cChar), cChar)
	var matrix [][]int //申请一个i行j列的矩阵
	matrix = make([][]int, iRow)
	//设置哨兵,
	matrix[0] = make([]int, jCol)
	for row := 1; row < iRow; row ++ {
		matrix[row] = make([]int, jCol)
		for col := 1; col < jCol; col ++ {
			if rChar[row] == cChar[col] {//i行j列相等时则左上方单格的值加1.
				matrix[row][col] = matrix[row-1][col-1] + 1
			} else {//i行j列不相等时,则选择上方和左方邻居中较大的那个.
				matrix[row][col] = max(matrix[row][col-1], matrix[row-1][col]) //区别于求公共子串.
			}
		}
	}
	pArr := make([]byte, 0)//存储公共字序列
	count := 0
	for _, rows := range matrix {
		for key, val := range rows {
			if val > count {
				count = val
				pArr = append(pArr, cChar[key])
			}
		}
	}
	//获取矩阵里最大的值.
	printSubSequence(rChar, cChar, matrix)
	fmt.Printf("%s与%s的公共子序列:%c\n", a, b, pArr)
	return string(pArr)
}
//打印矩阵,显示子串的关系
func printSubSequence(rowChar, colChar []byte, mtr [][]int) {
	for r, rows := range mtr {
		if r == 0 {
			continue
		}
		if r == 1 {
			fmt.Print(" ")
			for _, rc := range colChar {
				fmt.Print(string(rc), " ")
			}
			fmt.Println()
		}
		for c, col := range rows {
			if c == 0 {
				continue
			}
			if c == 1 {
				fmt.Printf("%c ", rowChar[r])
			}
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}
