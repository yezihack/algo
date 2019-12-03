package _8_哨兵的应用

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)


var arr = []int{1, 5, 3, 8, 9, 10}
func TestFind(t *testing.T) {
	src.Asset(1, t, 3, Find(arr, 8))
	src.Asset(2, t, 4, Find(arr, 9))
	src.Asset(3, t, 5, Find(arr, 10))
	src.Asset(4, t, -1, Find(arr, 100))
}
func TestFindBySentry(t *testing.T) {
	src.Asset(1, t, 3, FindBySentry(arr, 8))
	src.Asset(2, t, 4, FindBySentry(arr, 9))
	src.Asset(3, t, 5, FindBySentry(arr, 10))
	src.Asset(4, t, -1, FindBySentry(arr, 100))
}