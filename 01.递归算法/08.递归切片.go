package Recurive


//将二维数组递归转换成一维数组
func GetOne(matrix [][]int) []int {
	arr := make([]int, 0)
	//递归终止条件,即矩阵长度为0
	if len(matrix) == 0 {
		return arr
	}
	//将
	arr = append(matrix[0], GetOne(matrix[1:])...)
	return arr
}
