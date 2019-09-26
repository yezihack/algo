package leet_code_test

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

//思路: 利用相减的结果,错位存储map, 也就是说把相减的结果做key存入map里,待下次相减的结果在map里找到,并返回
func FindTwoSum(list []int, target int) (int, int) {
	checked := make(map[int]int)
	for i := 0; i < len(list); i++ {
		diff := target - list[i]        //相减获取差值
		if _, ok := checked[diff]; ok { //判断是否存在map里
			return checked[diff], i //返回map的值,与当前i为下标
		}
		checked[list[i]] = i //将相减的结果做key存储.
	}
	return -1, -1
}
