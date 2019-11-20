package LeetCode

import (
	"fmt"
	"sort"
)

//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//你可以假设数组是非空的，并且给定的数组总是存在众数。
func MajorityElement(nums []int) int {
	set := make(map[int]int)
	for _, v := range nums {
		if _, ok := set[v]; ok {
			set[v] += 1
		} else {
			set[v] = 1
		}
	}
	fmt.Println(set, len(nums)/2)
	for num, count := range set {
		if count > len(nums)/2 {
			return num
		}
	}
	return 0
}
func MajorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}