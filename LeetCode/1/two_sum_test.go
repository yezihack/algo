package leet_code

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/two-sum/

//题目意思是: 给定一个目标值,找出数组里2个数相加等于目标值的下标
func TestFindTwoSum(t *testing.T) {
	list := []int{2, 3, 7, 8}
	a, b := FindTwoSum(list, 15)
	fmt.Println(a, b)
}
func TestTwoSum(t *testing.T) {
	lst := []int{3,2,4}
	s := TwoSum(lst, 6)
	fmt.Println(s)
}
func TestTwoSum2(t *testing.T) {
	lst := []int{3,2,4,8,9,10}
	fmt.Println(TwoSum2(lst, 19))
	fmt.Println(TwoSum2(lst, 5))
	fmt.Println(TwoSum2(lst, 12))
	fmt.Println(TwoSum2(lst, 18))
	fmt.Println(TwoSum2(lst, 14))
}
func TestTwoSum3(t *testing.T) {
	lst := []int{3,2,4,8,9,10}
	fmt.Println(TwoSum3(lst, 19))
	fmt.Println(TwoSum3(lst, 5))
	fmt.Println(TwoSum3(lst, 12))
	fmt.Println(TwoSum3(lst, 18))
	fmt.Println(TwoSum3(lst, 14))
}

