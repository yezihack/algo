package Recurive

import (
	"fmt"
)

//利用i,j指针与一个基准值对比,将小于基准值的交互到最左边, 大小交互到最右边
//将后将最左,最右的数据不断排序,缩小范围.得到一个有序数组.
func QSort(a []int, left, right int) {
	pivot := a[left]
	i, j := left, right
	fmt.Printf("1.i:%d, j:%d, left:%d, right:%d, pivot:%d, %v\n", i, j, left, right, pivot, a)
	for i < j {
		for i < j && a[j] > pivot {
			j--
		}
		a[i] = a[j]
		for i < j && a[i] < pivot {
			i++
		}
		a[j] = a[i]
	}
	//i,j指针相偶
	a[i] = pivot
	fmt.Printf("2.i:%d, j:%d, left:%d, right:%d, pivot:%d, %v\n", i, j, left, right, pivot, a)
	if i-left > 1 {
		QSort(a, left, i-1)
	}
	if right-i > 1 {
		QSort(a, i+1, right)
	}
}
func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}
func quickSort(a []int, start, end int)  {
	if start < end {
		index := partition(a, start, end)
		quickSort(a, start, index-1)
		quickSort(a, index+1, end)
	}
}
func partition(a []int, left, right int) int {
	pivot := a[left]
	for left < right {
		for left < right && a[right] > pivot {
			right --
		}
		a[left] = a[right]
		for left < right && a[left] < pivot {
			left ++
		}
		a[right] = a[left]
	}
	a[left] = pivot
	return left
}
