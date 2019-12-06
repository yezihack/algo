package _9_查找算法

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7}
	src.Asset(1, t, 3, BinarySearch(arr, 4))
	src.Asset(2, t, 0, BinarySearch(arr, 1))
	src.Asset(3, t, 6, BinarySearch(arr, 7))
	src.Asset(4, t, -1, BinarySearch(arr, 70))
	src.Asset(5, t, -1, BinarySearch(arr, -50))
}
func TestBReserveSearch(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7}
	src.Asset(1, t, 3, BReserveSearch(arr, 0, len(arr)-1, 4))
	src.Asset(2, t, 0, BReserveSearch(arr, 0, len(arr)-1,1))
	src.Asset(3, t, 6, BReserveSearch(arr, 0, len(arr)-1,7))
	src.Asset(4, t, -1, BReserveSearch(arr, 0, len(arr)-1,70))
	src.Asset(5, t, -1, BReserveSearch(arr, 0, len(arr)-1,-50))
}
func TestBSqrt(t *testing.T) {
	var x float64 = 5
	fmt.Println(BSqrt(x))
}
func TestFirstSearch(t *testing.T) {
	arr := []int{1,2,3,3, 3,4,5}
	src.Asset(1, t, 2, BFirstSearch(arr, 3))
}
func TestBLastSearch(t *testing.T) {
	arr := []int{1,2,3,3,4,5}
	src.Asset(1, t, 3, BLastSearch(arr, 3))
}
func TestBGtSearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	src.Asset(1, t, 2, BGtSearch(arr, 4))
	src.Asset(2, t, 1, BGtSearch(arr, 2))
	src.Asset(3, t, 4, BGtSearch(arr, 9))
	src.Asset(4, t, 0, BGtSearch(arr, 1))
	src.Asset(5, t, 4, BGtSearch(arr, 10))
}
func TestBLTSearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	src.Asset(1, t, 1, BLTSearch(arr, 4))
	src.Asset(2, t, 0, BLTSearch(arr, 1))
	src.Asset(3, t, 4, BLTSearch(arr, 9))
	src.Asset(4, t, 3, BLTSearch(arr, 8))
	src.Asset(5, t, 4, BLTSearch(arr, 81))
}