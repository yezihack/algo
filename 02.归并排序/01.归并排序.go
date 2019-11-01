package main

import (
	"fmt"
)

func mergeSort(arr, temp []int, left, mid, right int) {
	i := left    //数组左边的起始位置
	j := mid + 1 //数组右边的起始位置
	t := 0       //临时数据的起始位置
	fmt.Printf("初使化: i:%d, j:%d, t:%d, left:%d, right: %d\n", i, j, t, left, right)
	//1. 把左右两边的有序数列按照规则填充到temp数组里
	//直到左右两边有序, 有一边处理完毕为止
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[t] = arr[i]
			i++
			t++
		} else {
			temp[t] = arr[j]
			j++
			t++
		}
	}
	//fmt.Printf("移动后:i:%d, j:%d, t:%d, left:%d, right: %d\n", i, j, t, left, right)
	//2. 将左边余下填充到temp数组里
	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}
	//把右边剩余的填充到temp数组里.
	for j <= right {
		temp[t] = arr[j]
		t++
		j++
	}
	//3. 将temp复制到arr
	tempLeft := left
	t = 0
	for tempLeft <= right {
		arr[tempLeft] = temp[t]
		t++
		tempLeft++
	}
	//	fmt.Printf("%s", debug.Stack())
}
func merge(arr, temp []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		merge(arr, temp, left, mid)
		merge(arr, temp, mid+1, right)
		//fmt.Printf("left:%d, mid: %d, right:%d\n", left, mid, right)
		mergeSort(arr, temp, left, mid, right)
	}
}

//参考视频: https://www.bilibili.com/video/av54029771/?p=70
func main() {
	a := []int{1, 3, 5, 6, 8, 7, 2, 9}
	fmt.Println("排序前", a)
	result := make([]int, len(a))
	merge(a, result, 0, len(a)-1)
	fmt.Println("排序后", result)
}
