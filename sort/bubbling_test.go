package sort

import (
	"fmt"
	"testing"
)

//冒泡算法O(n^2)
//双循环, 每次拿一个值与数组其它值相比较, 如果i小于j则交互位置
func bubbling(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				count++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("count", count)
	return arr
}

func TestBubbling(t *testing.T) {
	a := []int{4, 2, 3, 1}
	fmt.Println(bubbling(a))
}

//升级版本
func Bubbing2(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			flag = true
		}
		if !flag {
			break
		}
	}
	return arr
}
func TestBubbling2(t *testing.T) {
	a := []int{1, 8, 3, 2, 9, 4}
	fmt.Println(Bubbing2(a))
}

func Test3(t *testing.T) {
	a, b := -540, -33
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a", a, "b", b)
}
