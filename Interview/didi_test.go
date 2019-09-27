package Interview

import (
	"fmt"
	"testing"
)

// 滴滴面试的算法题

//1. 一个矩阵,打印成蛇形
//以下为矩阵
// 1. 2. 3. 4. 5
// 6. 7. 8. 9. 10
// 11.12.13.14.15
// 16.17.18.19.20
// 需要打印成
//1.2.3.4.5 10.15.20.19.18.17.16.11.6.7.8.9.14.13.12

func MatrixPrintSnake(Matrix [][]int) string {
	return ""
}

//2. 手写一个快速排序
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := arr[0]
	less := make([]int, 0)
	great := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if mid > arr[i] {
			less = append(less, arr[i])
		} else {
			great = append(great, arr[i])
		}
	}
	fmt.Println("less", less)
	fmt.Println("great", great)
	result := make([]int, 0)
	result = append(result, QuickSort(less)...)
	result = append(result, mid)
	result = append(result, QuickSort(great)...)
	return result
}
func TestQuickSort(t *testing.T) {
	a := []int{8, 2, 4}
	fmt.Println(QuickSort(a))
}
