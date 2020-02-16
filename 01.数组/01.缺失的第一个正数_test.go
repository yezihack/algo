package _1_数组

import (
	"fmt"
	"testing"
)

/*
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
说明:

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//解法一: 利用map+for-range实现.时间复杂度O(n), 空间是常数O(n)
//	4 ms	2.8 MB
func FirstMissingPositive(nums []int) int {
	hash := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i ++ {
		hash[nums[i]] = struct{}{}
	}
	fmt.Println(hash)
	//1-n之间检查, 如果有缺失则是最小值.
	for i := 1; i <= len(nums); i ++ { //从1循环到n, 包含n
		if _,ok := hash[i]; !ok {
			return i
		}
	}
	//如果都不在hash里面的话,则是长度+1
	return len(nums) + 1
}
func TestFirstMissingPositive(t *testing.T) {
	tests := []struct{
		index int //序号
		data []int //数据
		expect int //预期值
	}{
		{1, []int{1, -1, 3, 4}, 2},
		{2, []int{7,8,9,11,12}, 1},
		{3, []int{1,2,0}, 3},
		{4, []int{3,4,-1,1}, 2},
	}
	for _, item := range tests {
		if actual := FirstMissingPositive(item.data); actual != item.expect {
			t.Errorf("index:%d, expect:%d, actual:%d\n", item.index, item.expect, actual)
		}
	}
}

//原地桶排序一次. 原地交换. 如1则放在0的位置.
func FirstMissingPositiveV2(nums []int) int {
	length := len(nums)
	fmt.Println(nums)
	for i := 0; i < length; i ++{
		for nums[i] > 0 && nums[i] <= length && nums[i] != nums[nums[i]-1]{
			tmp  := nums[i]
			fmt.Printf("tmp:%d, i:%d, i+1:%d\n", tmp, nums[i], nums[i+1])
			nums[i] = nums[nums[i]-1]
			nums[nums[i]-1] = tmp
		}
	}
	for i := 0; i <= length; i ++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return length+1
}
func TestFirstMissingPositiveV2(t *testing.T) {
	tests := []struct{
		index int //序号
		data []int //数据
		expect int //预期值
	}{
		{1, []int{1, -1, 3, 4}, 2},
		{2, []int{7,8,9,11,12}, 1},
		{3, []int{1,2,0}, 3},
		{4, []int{3,4,-1,1}, 2},
	}
	for _, item := range tests {
		if actual := FirstMissingPositiveV2(item.data); actual != item.expect {
			t.Errorf("index:%d, expect:%d, actual:%d\n", item.index, item.expect, actual)
		}
	}
}