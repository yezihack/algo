package _3_动态规划

import "fmt"

//计算fish和fosh的最长公共子串序列
//先抽象成数据模型. 利用矩阵查找关系
//思路: 先目标字符串抽象成字符行, 匹配字符串抽象成字符列. 行与列就成了矩阵
// 若row为行, col为列,比较之间的关系.
// 如果相等则选择左上方单元格的值加1
// 如果不相等则选择左边相邻和上方单元格中较大的那个值.
func PubStr(a, b string) string {
	rChar := []byte(a)
	cChar := []byte(b)
	iRow, jCol := len(rChar), len(cChar)
	var matrix [][]int //申请一个i行j列的矩阵
	matrix = make([][]int, iRow)
	for row := 0; row < iRow; row ++ {
		matrix[row] = make([]int, jCol)
		for col := 0; col < jCol; col ++ {
			if row == 0 && col == 0 {//第一个单元格,特殊处理
				if rChar[row] == cChar[col] {//如果两者相等则赋值1
					matrix[row][col] = 1
				} else {//如果不相等则赋值0
					matrix[row][col] = 0
				}
			} else if rChar[row] == cChar[col] {//i行j列相等时则左上方单格的值加1.
				matrix[row][col] = matrix[row-1][col-1] + 1
			} else {//i行j列不相等时,则选择上方和左方邻居中较大的那个.
				if row == 0 {
					matrix[row][col] =  matrix[row][col-1]
				} else if col == 0 {
					matrix[row][col] =  matrix[row-1][col]
				} else {
					matrix[row][col] = max(matrix[row][col-1], matrix[row-1][col])
				}
			}
		}
	}
	pArr := make([]byte, 0)//存储公共字符
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
	printMatrix(rChar, cChar, matrix)
	fmt.Printf("%s与%s的公共字符串:%c\n", a, b, pArr)
	return string(pArr)
}
//打印矩阵,显示子串的关系
func printMatrix(rowChar, colChar []byte, mtr [][]int) {
	for r, rows := range mtr {
		if r == 0 {
			fmt.Print("  ")
			for _, rc := range colChar {
				fmt.Print(string(rc), " ")
			}
			fmt.Println()
		}
		for c, col := range rows {
			if c == 0 {
				fmt.Printf("%c ", rowChar[r])
			}
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
	fmt.Println()
}
