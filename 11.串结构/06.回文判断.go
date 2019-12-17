package string_string

import "fmt"

//回文判断, 使用两个指针,向中间逼近
func Palindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) < 2 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left ++
		right --
	}
	return true
}

//解题思路:
// 与左右指针向中间逼近的思路刚好相反的思路
// 先从中间开始向两边扩散逐一比较字符
func PalindromeByMid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) < 2 {
		return true
	}
	mid := len(s) >> 1
	var left, right int
	if len(s) % 2 == 0 {
		left, right = mid-1, mid
	} else {
		left, right = mid, mid
	}
	fmt.Printf("%s, mid:%d, left:%d, right:%d\n", s, mid, left, right)
	for left >= 0 && right <= len(s)-1 {
		if s[left] != s[right] {
			return false
		}
		left--
		right++
	}
	return true
}
