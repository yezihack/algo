package _2_数学

import (
	"fmt"
	"strconv"
	"testing"
)

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//方法一: 转称字符串进行判断.
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i ++
		j --
	}
	return true
}
func TestIisPalindrome(t *testing.T) {
	tests := []struct {
		index int
		data int
		expect bool
	}{
		{1, 121, true},
		{2, -121, false},
		{3, 10, false},
	}
	for _, item := range tests {
		if ret := isPalindrome(item.data); ret != item.expect {
			t.Errorf("expect:%v, actual:%v\n", item.expect, ret)
		}
	}
}

//方法2, 求mod, 得到数字的反转结果. 因为回文数字,与原数字是相等的.
//但是有溢出的风险.
func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}
	y := 0
	xx := x
	for xx != 0 {
		mod := xx % 10
		xx /= 10
		y = y * 10 + mod
	}
	fmt.Printf("x:%d, y:%d\n", x, y)

	if x == y {
		return true
	}
	return false
}
func TestIisPalindromeV2(t *testing.T) {
	tests := []struct {
		index int
		data int
		expect bool
	}{
		{1, 121, true},
		{2, -121, false},
		{3, 10, false},
	}
	for _, item := range tests {
		if ret := isPalindromeV2(item.data); ret != item.expect {
			t.Errorf("index:%d, expect:%v, actual:%v\n",item.index, item.expect, ret)
		}
	}
}
func isPalindromeV3(x int) bool {
	if x < 0 || (x != 0 && x % 10 == 0){
		return false
	}
	y := 0
	for x > y {
		y = y * 10 + x % 10
		x /= 10
	}
	fmt.Printf("x:%d, y:%d\n", x, y)
	if x == y || x == y/10 {
		return true
	}
	return false
}
func TestIisPalindromeV3(t *testing.T) {
	tests := []struct {
		index int
		data int
		expect bool
	}{
		{1, 121, true},
		{2, -121, false},
		{3, 10, false},
	}
	for _, item := range tests {
		if ret := isPalindromeV3(item.data); ret != item.expect {
			t.Errorf("index:%d, expect:%v, actual:%v\n",item.index, item.expect, ret)
		}
	}
}