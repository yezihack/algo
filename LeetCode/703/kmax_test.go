package _03

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	nums := []int{8, 9, 1, 3, 2, 7}
	k := Constructor(3, nums)
	fmt.Println(k)
	src.So(t, 7, k.Add(2))
	src.So(t, 8, k.Add(10))
	src.So(t, 9, k.Add(12))
	src.So(t, 10, k.Add(13))
	src.So(t, 10, k.Add(3))
}
