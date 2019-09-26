package sort

import (
	"fmt"
	"testing"
)

//冒泡算法O(n^2)
func bubbling(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func TestBubbling(t *testing.T) {
	a := []int{1, 8, 3, 2, 9, 4}
	fmt.Println(bubbling(a))
}
