package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

//快排
//先找基准值
//然后分区, 以基准值为准, 找出小于基准值与大于基准值的元素.
//空数组和一个元素的数组都无需排序.
func quickSort(arr []int) []int {
	if len(arr) < 2 { //空数组和一个元素的数组都无需排序
		return arr
	}
	//找基准值
	pivot := arr[0]
	//分区,归纳
	//查找小于基准值的元素
	less := make([]int, 0)
	//查找大于基准值的元素
	greater := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}
	//合并一个新数组
	newArr := make([]int, 0)
	newArr = append(newArr, quickSort(less)...)
	newArr = append(newArr, pivot)
	newArr = append(newArr, quickSort(greater)...)
	return newArr
}
func TestQuickSort(t *testing.T) {
	arr := []int{9, 2, 3, 21, 0, 1, 4}
	fmt.Println(quickSort(arr))
}
func TestQuickSort2(t *testing.T) {
	arr := []int{8, 2, 4}
	fmt.Println(quickSort(arr))
}

func QuickSortCommon(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[rand.Intn(len(arr))]
	less, greater, newArr := make([]int, 0), make([]int, 0), make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}
	return newArr
}
func ChildQuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	return nil
}
