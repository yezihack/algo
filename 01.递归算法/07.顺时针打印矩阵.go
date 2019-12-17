package Recurive

import "fmt"

//一个二维数组, 使用顺时针打印,类似蛇形.
/*
[
[1, 2, 3]
[4, 5, 6]
[7, 8, 9]
]

输出[1, 2, 3, 6, 9, 8, 7, 4, 5]
这道是在剑指offer里29题有讲.参数csdn:https://blog.csdn.net/hefenglian/article/details/78339990

解题思路1: 使用递归思维
如果一个更大的二维数据, 输出最外层,里面又是一个完整二维数据,就可以使用递归实现
第一步：找出递归出口，就是当只剩一个元素，即是出口．
第二步：写出递推公式:
f(n)=f(child(n-2))
child(n)=a[row-2][col-2]
1. 输出第一行,从头输到尾.
2. 输出最后一列, 排除第一行的最后一列已经输出了.
3. 输出最后一行, 排除最后一行的最后一列已经输出了.
4. 输出第一列, 需要排除第一行和最后一行已经输出了.
5. 然后将剩余的元素再次构建一个子二维数据,进行递归.

解题思路二: 非递归思维,实现: PrintMatrix
我们可以找一个最左上解作为起点, 如第一个圈起点坐标是(0,0), 第二个圈起点坐标是(1,1),以此类拴.
需要注意边界条件
分三种情况:
1. 4*4的矩阵,最后只有4个元素,只能打印从左到右, 从上到下,从右到左,而没有从下到上的打印
1. 4*3的矩阵,最后只有2个元素,只能打印从左到右. 没有从上到下, 从右到左, 从下到上的打印
1. 3*4的矩阵,最后只有2个元素,只能打印从左到右,从上到下,没有从右到左,从下到上的打印.
*/

// 递归实现
// 顺时针打印数组
func ClockWisePrint(arr [][]int) []int {
	result := make([]int, 0)
	//求出行与列的大小
	row, col := len(arr), len(arr[0])
	//写出递归出口
	fmt.Printf("arr:%v, row:%d, col:%d\n", arr, row, col)
	if row == 1 && col == 1 {
		result = append(result, arr[row-1][col-1])
		return result
	}
	//打印第一行
	for i := 0; i < col; i ++ {
		fmt.Println("第一行", arr[0][i])
		result = append(result, arr[0][i])
	}
	//打印最后一列
	if col >= 1 {
		for i := 1; i < row; i ++ {
			fmt.Println("最后一列:", arr[i][col-1])
			result = append(result, arr[i][col-1])
		}
	}
	//打印最后一行
	if col > 2 && row >= 2 {
		for i := col-2; i >= 0; i-- {
			fmt.Println("最后一行:", arr[row-1][i])
			result = append(result, arr[row-1][i])
		}
	}
	//打印第一列
	if row > 2 && col > 2 {
		for i := row-2; i > 0; i-- {
			fmt.Println("第一列:", arr[i][0])
			result = append(result, arr[i][0])
		}
	}
	//将剩余的元素再构造一个子二维数组, 递归调用.
	if row - 2 > 0 && col - 2 > 0{
		//初使一个子二维切片里
		child := make([][]int, row-2)
		for i := 0; i < row-2;i ++ {
			child[i] = make([]int, col-2)
		}
		//去掉外围已经打印的元素,所剩余的元素装载到一个新的子二维切片里
		for i := 1; i <= row-2;i ++ {
			for j := 1; j <=col-2; j ++ {
				child[i-1][j-1] = arr[i][j]
			}
		}
		fmt.Println("child", child)
		result = append(result, ClockWisePrint(child)...)
	}
	return result
}

//解题3思路: 每一次从最左上角打印,即坐标为(0,0),打印最外层一圈. 然后打印第二圈,即坐标:(1,1), 以次类推
//时间复杂度:O(n)
func PrintMatrixLikeSnake(matrix [][]int) []int {
	if len(matrix) <= 0 {
		return nil
	}
	rows := len(matrix)
	columns := len(matrix[0])
	if columns <= 0 {
		return nil
	}
	//定义一个结果集
	result := make([]int, 0)
	//准备打印第一个最外层的圈,定义一个启始坐标
	start := 0
	//我们发现5*5的矩阵, 最后一个圈坐标为(2,2), 4*4矩阵最后一个圈的坐标是(1,1),得出退出条件X坐标大于2倍起点坐标.Y坐标同理可得.
	for rows > start * 2 && columns > start * 2 {
		endX := columns-1-start
		endY := rows-1-start
		//从左到右打印,即第一行的边
		for i := start; i <= endX; i++ {
			result = append(result, matrix[start][i])
		}
		//从最右边从上到下打印,即最右边的边
		if start < endY {
			for i := start+1;i <= endY; i ++ {
				result = append(result, matrix[i][endX])
			}
		}
		//从最右边从右到左打印,即最下面的边
		if start < endX && start < endY {
			for i := endX-1; i >= start; i -- {
				result = append(result, matrix[endY][i])
			}
		}
		//从最左边从下到上打印,即最左边的边
		if start < endY - 1 {
			for i := endY-1; i >= start+1; i -- {
				result = append(result, matrix[i][start])
			}
		}
		start ++
	}
	return result
}

