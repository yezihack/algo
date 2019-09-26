package bianry_search

import (
	"testing"
)

//二分查找
//数组必须是有序的.否则免谈
//每一次查找都是折半查找. 比对目标值, 四种情况, 如下:
//第一种情况:相等则返回下标
//第二种情况: 目标值小于中间值, 则修改high值,即high - 1
//第三种情况: 目标值大于中间值, 则修改low值, 即low + 1
//第四种情况: 未找到.返回 -1
func binarySearch(arr []int, search int) int {
	low := 0             //定义最低位
	high := len(arr) - 1 //定义最高位, 即数组的最后的下标
	for low <= high {    //缩小到只包含一个元素
		mid := (low + high) / 2
		if search == arr[mid] {
			return mid
		} else if search < arr[mid] {
			high = mid - 1
		} else if search > arr[mid] {
			low = mid + 1
		} else {
			return -1
		}
	}
	return -1
}
func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 7, 9, 10, 22, 88}
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			"二分查找1",
			1,
			0,
		},
		{
			"二分查找2",
			22,
			5,
		},
		{
			"二分查找3",
			100,
			-1,
		},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			result := binarySearch(arr, item.args)
			if result != item.want {
				t.Errorf("want: %d, err:%d\n", item.want, result)
			}
		})
	}
}
