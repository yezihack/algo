package sort

import (
	"fmt"
	"testing"
)

//选择排序
//思路:　每一次查找每选择一个最小值或最大值．然后放在一个新的数组里．
//第一个循环是从左到右
//第二个循环是从右到左
//将第2个循环与第一个循环相比, 谁最小就交互位置.
func selectSort(arr []int) []int {
	lenth := len(arr)
	for i := 0; i < lenth; i++ {
		small := i                       //默认将最小值从0开始.
		for j := lenth - 1; j > i; j-- { //从最右与i相比
			if arr[j] < arr[small] {
				small = j
			}
		}
		//查找到的最小值与之前的位置交互
		arr[i], arr[small] = arr[small], arr[i]
	}
	return arr
}
func TestSelectSort(t *testing.T) {
	arr := []int{1, 8, 2, 3, 3, 10, 9, 4}
	newArr := selectSort(arr)
	fmt.Println(newArr)
}
