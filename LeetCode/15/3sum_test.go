package leet_code

import (
	"fmt"
	"testing"
)
//暴力方法
func TestThreeSum(t *testing.T) {
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(ThreeSum([]int{0, 0, 0, 0}))
	fmt.Println(ThreeSum([]int{-2, 0, 0, 2, 2}))
	fmt.Println(ThreeSum([]int{3, 0, -2, -1, 1, 2}))
}
func BenchmarkThreeSum(b *testing.B) {
	for i := 0; i < b.N; i ++{
		ThreeSum([]int{3,0,-2,-1,1,2})
		//500000	      2816 ns/op
	}
}

//二个循环,借用map法
func TestThreeSumMap(t *testing.T) {
	fmt.Println(ThreeSumMap([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(ThreeSumMap([]int{0,0,0,0}))
	fmt.Println(ThreeSumMap([]int{-2,0,0,2,2}))
}
//左右指针法
func TestThreeSumPointer(t *testing.T) {
	fmt.Println(ThreeSumPointer([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(ThreeSumPointer([]int{0,0,0,0}))
	fmt.Println(ThreeSumPointer([]int{-2,0,0,2,2}))
}



