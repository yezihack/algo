package _3_动态规划

import (
	"fmt"
	"testing"
)
// 一道面试题
// 问题是这样的：一个整数数组中的元素有正有负，在该数组中找出一个连续子数组，要求该连续子数组中各元素的和最大，
// 这个连续子数组便被称作最大连续子数组。比如数组{3, -1, 5, -8, 3, 3}的最大连续子数组为{3, -1, 5}，
// 最大连续子数组的和为3-1+5=7


//设置一个求解最大值的变量max, 然后从头到尾计算一步, 与max比较一下, 最终求得最大值.
//复杂度O(n^2)
func SearchSum(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i ++ {
		sum := 0
		for j := i; j < len(arr); j ++ {
			sum += arr[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func TestSearchSum(t *testing.T) {
	fmt.Println(SearchSum([]int{3, -1, 5, -8, 3, 3}))
}
//题目明确, 有正有负, 最大值, 一定大于0.
//从头加到尾,每加一步都与最大值比较一下, 如果小于0则重新计算.
//属于动态规划思想.
func SearchSumV2(arr []int) int {
	sum := 0 //每次累加的变量值
	max := 0 //最终求得的最大值.
	for i := 0; i < len(arr); i ++{
		sum += arr[i] //累加值
		if sum > max {//每计算一步就比较之前的max值.
			max = sum
		}
		//如果sum小于0则重新计算
		if sum <= 0 {
			sum = 0
		}
	}
	return max
}

func TestCalcJoinMax(t *testing.T) {
	fmt.Println(SearchSumV2([]int{3, -1, 5, -8, 3, 3}))
}