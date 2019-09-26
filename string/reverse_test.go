package _string

import (
	"fmt"
	"testing"
)

//反转字符串
//从数组最后一位开始倒着读
func reverse(s []byte) []byte {
	r := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
func TestReverse(t *testing.T) {
	s := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println(string(reverse(s)))
}

//反转字符串2
//使用位置交换实现倒序
func reverse2(s string) string {
	r := []rune(s)
	for from, to := 0, len(r)-1; from < to; from, to = from+1, to-1 {
		r[from], r[to] = r[to], r[from
	}]
	return string(r)
}
func TestReverse2(t *testing.T) {
	s := "abcdef"
	fmt.Println(reverse2(s))
}
