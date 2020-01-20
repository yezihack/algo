package _6_排序算法

import (
	"fmt"
	"testing"
)

func TestQuickSort4(t *testing.T) {
	arr := []int{8, 2, 1, 2, 9, 7}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func TestQuickSort5(t *testing.T) {
	arr := []int{8, 2, 1, 2, 9, 7}
	QuickSortSimple(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func TestQuickSortOne(t *testing.T) {
	arr := []int{6, 2, 8, 1, 3, 4, 7, 5, 9}
	QuickSortOne(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arr = []int{6, 2, 8, 2, 3, 4, 5, 5, 1}
	QuickSortOne(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func TestQuickSort6(t *testing.T) {
	arr := []int{6, 2, 8, 1, 3, 4, 7, 5, 9}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}