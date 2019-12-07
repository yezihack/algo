package _3_动态规划

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestPubStr(t *testing.T) {
	src.Asset(1, t, "lue", SubSequence("clues","blue"))
	src.Asset(2, t, "fsh", SubSequence("fish","fosh"))
	src.Asset(3, t, "llo", SubSequence("hello","kallo"))
	src.Asset(4, t, "ks", SubSequence("kiss","kas"))
}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3}
	s1 := src.NewInt().Clone(s)
	fmt.Println(InsertIndex(s, 0, 0))

	fmt.Println(src.NewInt().InsertIndex(s1, 2, 5))
}
func InsertIndex(arr []int, index, value int) []int  {
	rear := append([]int{}, arr[index:]...)
	return append(append(arr[:index], value), rear...)
}
