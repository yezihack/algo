package LeetCode

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//暴力求解.O(n^2)
func TwoSumPower(nums []int, target int) []int{
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if k1 != k2 && target == v1 + v2 {
				return []int{k1, k2}
			}
		}
	}
	return []int{0,0}
}

func TestTwoSumPower(t *testing.T) {
	src.Asset(1, t, []int{0, 1}, TwoSumPower([]int{2, 7, 11, 15}, 9))
	src.Asset(2, t, []int{0, 1}, TwoSumPower([]int{3, 3}, 6))
}
//利用一个map, 求一个差值. 时间O(n), 空间O(n)
func twoSum(nums []int, target int) []int {
	set := make(map[int]int, len(nums))
	for k, v := range nums {
		if k1, ok := set[target - v]; ok {
			return []int{k1, k}
		}
		set[v] = k
	}
	return []int{0,0}

}
func TestTwoSum(t *testing.T) {
	src.Asset(1, t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	src.Asset(2, t, []int{0, 1}, twoSum([]int{3, 3}, 6))
}