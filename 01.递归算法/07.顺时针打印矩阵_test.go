package Recurive

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestPrintMatrixLikeSnake(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := []int{1,2,3,6,9,8,7,4,5}
	actual := PrintMatrixLikeSnake(a)
	src.Asset(1, t, expect, actual)
}
func TestPrintMatrixLikeSnake2(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{11, 22, 33},
	}
	expect := []int{1,2,3,6,9,33,22,11,7,4,5,8}
	actual := PrintMatrixLikeSnake(a)
	src.Asset(1, t, expect, actual)
}
func TestPrintMatrixLikeSnake3(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
	}
	expect := []int{1,2,3}
	actual := PrintMatrixLikeSnake(a)
	src.Asset(1, t, expect, actual)
}
func TestPrintMatrixLikeSnake4(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	expect := []int{1,2,3,6,5,4}
	actual := PrintMatrixLikeSnake(a)
	src.Asset(1, t, expect, actual)
}
func TestPrintMatrixLikeSnake5(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	expect := []int{1,2,3,4,7,11,10,9,8,4,5,6}
	actual := PrintMatrixLikeSnake(a)
	src.Asset(1, t, expect, actual)
}
func TestPrintMatrixLikeSnake6(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 51},
		{1, 2, 99, 4, 52},
		{1, 2, 3, 4, 53},
		{1, 2, 3, 4, 54},
	}
	expect := []int{1,2,3,4,5,51,52,53,54,4,3,2,1,1,1,1,2,3,4,4,4,3,2,2,99}
	actual := PrintMatrixLikeSnake(a)
	src.Asset(1, t, expect, actual)
}
func TestClockWisePrint(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := []int{1,2,3,6,9,8,7,4,5}
	actual := ClockWisePrint(a)
	src.Asset(1, t, expect, actual)
}
func TestClockWisePrint2(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{11, 22, 33},
	}
	expect := []int{1,2,3,6,9,33,22,11,7,4,5,8}
	actual := ClockWisePrint(a)
	src.Asset(1, t, expect, actual)
}
func TestClockWisePrint3(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
	}
	expect := []int{1,2,3}
	actual := ClockWisePrint(a)
	src.Asset(1, t, expect, actual)
}
func TestClockWisePrint4(t *testing.T) {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	expect := []int{1,2,3,6,5,4}
	actual := ClockWisePrint(a)
	src.Asset(1, t, expect, actual)
}
func TestClockWisePrint5(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	expect := []int{1,2,3,4,7,11,10,9,8,4,5,6}
	actual := ClockWisePrint(a)
	src.Asset(1, t, expect, actual)
}
func TestClockWisePrint6(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 51},
		{1, 2, 99, 4, 52},
		{1, 2, 3, 4, 53},
		{1, 2, 3, 4, 54},
	}
	expect := []int{1,2,3,4,5,51,52,53,54,4,3,2,1,1,1,1,2,3,4,4,4,3,2,2,99}
	actual := ClockWisePrint(a)
	src.Asset(1, t, expect, actual)
}
