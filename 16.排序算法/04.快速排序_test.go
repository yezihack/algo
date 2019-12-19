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
