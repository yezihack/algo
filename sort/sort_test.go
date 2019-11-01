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
func TestQsort(t *testing.T) {
	list := []int{5, 3, 4, 2, 1}
	QSort(list, 0, len(list)-1)
	if list[0] != 1 || list[1] != 2 || list[2] != 3 || list[3] != 4 || list[4] != 5 {
		t.Error("排序失败", list)
	}
	fmt.Println(list)
	list = []int{2, 3, 5, 4, 1}
	QSort(list, 0, len(list)-1)
	if list[0] != 1 || list[1] != 2 || list[2] != 3 || list[3] != 4 || list[4] != 5 {
		t.Error("排序失败", list)
	}
	fmt.Println(list)
	list = []int{5, 4}
	QSort(list, 0, len(list)-1)
	if list[0] != 4 || list[1] != 5 {
		t.Errorf("sort fail. list:%v", list)
	}
	fmt.Println(list)
	list = []int{5, 9, 2, 3, 1}
	QSort(list, 0, len(list)-1)
	if list[0] != 1 || list[1] != 2 || list[4] != 9 {
		t.Errorf("sort fail. list:%v", list)
	}
	fmt.Println(list)
}
