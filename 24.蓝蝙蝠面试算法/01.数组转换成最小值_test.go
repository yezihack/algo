package _4_蓝蝙蝠面试算法

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

/*
如arr := []int{16, 15, 27, 9}
输出: 1516279
*/
func TestPrintMinNumber(t *testing.T) {
	src.Asset(1, t, "10211516279", PrintMinNumberByBubbling([]int {16,27,15,9,102, 1}))
	src.Asset(2, t, "1516279", PrintMinNumberByBubbling([]int{16,15,27,9}))
}

func TestMinNumber(t *testing.T) {
	src.Asset(1, t, "10211516279", MinNumberBySort([]int {16,27,15,9,102, 1}))
	src.Asset(2, t, "1516279", MinNumberBySort([]int{16,15,27,9}))
}

func TestMinNumbersByQuickSort(t *testing.T) {
	src.Asset(1, t, "10211516279", MinNumbersByQuickSort([]int {16,27,15,9,102, 1}))
	src.Asset(2, t, "1516279", MinNumbersByQuickSort([]int{16,15,27,9}))
}

